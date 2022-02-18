package auth

import (
	"context"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func (f *Firebase) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	return f.Auth.VerifyIDToken(ctx, idToken)
}

func SetupFirebase() Firebase {

	env := os.Getenv("VOTODEV_JSON")
	new := strings.Replace(env, "`", "", -1)

	opt := option.WithCredentialsJSON([]byte(new))
	ctx := context.Background()
	//Firebase admin SDK initialization
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic("Firebase load error: " + err.Error())
	}
	//Firebase Auth
	auth, err := app.Auth(ctx)
	if err != nil {
		panic("Firebase load error: " + err.Error())
	}
	f := Firebase{
		Auth: auth,
		App:  app,
		Ctx:  ctx,
	}
	return f
}

func (f *Firebase) CreateAuthUser(email string, phone string, pwd string, name string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		Password(pwd).
		DisplayName(name).
		Disabled(false)
	u, err := f.Auth.CreateUser(context.Background(), params)
	return u, err
}

func (f *Firebase) DeleteAuthUser(email string) error {
	authuser, err := f.Auth.GetUserByEmail(context.Background(), email)
	if err != nil {
		return err
	}
	err = f.Auth.DeleteUser(context.Background(), authuser.UID)
	if err != nil {
		return err
	}
	return nil
}
