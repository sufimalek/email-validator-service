package server

import (
	"encoding/json"
	"net/http"

	"github.com/sufimalek/emailvalidator"
	"github.com/sufimalek/emailvalidator/validators"
)

type ValidationRequest struct {
	Email string `json:"email"`
}

type ValidationResponse struct {
	IsValid bool   `json:"is_valid"`
	Message string `json:"message,omitempty"`
}

func ValidateEmail(w http.ResponseWriter, r *http.Request) {
	var req ValidationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chain := emailvalidator.NewValidatorChainBuilder().
		Add(&validators.SyntaxValidator{}).
		Add(&validators.RoleValidator{}).
		Add(&validators.MXValidator{}).
		Add(&validators.SMTPValidator{}).
		Add(validators.NewBanWordsUsernameValidator([]string{"spam", "junk"})).
		Add(validators.NewBlackListEmailsValidator([]string{"blacklisted@example.com"})).
		Add(&validators.DNSCheckValidator{}).
		Add(&validators.RFCValidator{}).
		Add(&validators.NoRFCWarningsValidator{}).
		// Add(&validators.GravatarValidator{}).
		Build()

	isValid := chain.Validate(req.Email)

	response := ValidationResponse{
		IsValid: isValid,
	}
	if !isValid {
		response.Message = "Email validation failed"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
