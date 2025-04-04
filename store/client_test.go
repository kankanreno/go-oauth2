package store_test

import (
	"context"
	"testing"

	"github.com/kankanreno/go-oauth2/v4/models"
	"github.com/kankanreno/go-oauth2/v4/store"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClientStore(t *testing.T) {
	Convey("Test client store", t, func() {
		clientStore := store.NewClientStore()

		err := clientStore.Add(context.Background(), "1", &models.Client{ID: "1", Secret: "2"})
		So(err, ShouldBeNil)

		cli, err := clientStore.Get(context.Background(), "1")
		So(err, ShouldBeNil)
		So(cli.GetID(), ShouldEqual, "1")

		err = clientStore.Delete(context.Background(), "1")
		So(err, ShouldBeNil)
	})
}
