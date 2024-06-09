package server

import (
	"encoding/json"
	"net/http"

	"github.com/sufimalek/emailvalidator"
	"github.com/sufimalek/emailvalidator/validators"
)

type ValidateEmailRequest struct {
	Email string `json:"email"`
}

type ValidateEmailResponse struct {
	IsValid bool   `json:"is_valid"`
	Message string `json:"message"`
}

func ValidateEmailHandler(w http.ResponseWriter, r *http.Request) {
	var req ValidateEmailRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
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
	message := "Email is valid"
	if !isValid {
		message = "Email validation failed"
	}

	resp := ValidateEmailResponse{
		IsValid: isValid,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
