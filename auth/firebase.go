package auth

import (
	"context"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func (f *Firebase) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	return f.Auth.VerifyIDToken(ctx, idToken)
}

func SetupFirebase() Firebase {

	env := os.Getenv("VOTODEV_JSON")
	new := strings.Replace(env, "`", "", -1)

	opt := option.WithCredentialsJSON([]byte(new))
	logrus.Info("Created options")
	ctx := context.Background()
	logrus.Info("Created context")
	//Firebase admin SDK initialization
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		logrus.Error("Firebase load error: " + err.Error())
		panic("Firebase load error: " + err.Error())
	}
	//Firebase Auth
	logrus.Info("Created new app")
	auth, err := app.Auth(ctx)
	logrus.Info("Created new app auth")
	if err != nil {
		logrus.Error("Firebase app load error: " + err.Error())
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
	log.Println("Create Auth User in firebase")
	params := (&auth.UserToCreate{}).
		Email(email).
		EmailVerified(false).
		Password(pwd).
		DisplayName(name).
		Disabled(false)
	u, err := f.Auth.CreateUser(context.Background(), params)
	if err != nil {
		log.Println("Could not create auth user: ", err.Error())
	}
	return u, err
}

func (f *Firebase) DeleteAuthUser(email string) error {
	log.Println("Delete Auth User in firebase")
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
