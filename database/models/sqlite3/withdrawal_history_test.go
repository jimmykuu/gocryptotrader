// Code generated by SQLBoiler 3.5.0-gct (https://github.com/thrasher-corp/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlite3

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/thrasher-corp/sqlboiler/randomize"
	"github.com/thrasher-corp/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testWithdrawalHistories(t *testing.T) {
	t.Parallel()

	query := WithdrawalHistories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testWithdrawalHistoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WithdrawalHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWithdrawalHistoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := WithdrawalHistories().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WithdrawalHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWithdrawalHistoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WithdrawalHistorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WithdrawalHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWithdrawalHistoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := WithdrawalHistoryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if WithdrawalHistory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected WithdrawalHistoryExists to return true, but got false.")
	}
}

func testWithdrawalHistoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	withdrawalHistoryFound, err := FindWithdrawalHistory(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if withdrawalHistoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testWithdrawalHistoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = WithdrawalHistories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testWithdrawalHistoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := WithdrawalHistories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testWithdrawalHistoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	withdrawalHistoryOne := &WithdrawalHistory{}
	withdrawalHistoryTwo := &WithdrawalHistory{}
	if err = randomize.Struct(seed, withdrawalHistoryOne, withdrawalHistoryDBTypes, false, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, withdrawalHistoryTwo, withdrawalHistoryDBTypes, false, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = withdrawalHistoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = withdrawalHistoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := WithdrawalHistories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testWithdrawalHistoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	withdrawalHistoryOne := &WithdrawalHistory{}
	withdrawalHistoryTwo := &WithdrawalHistory{}
	if err = randomize.Struct(seed, withdrawalHistoryOne, withdrawalHistoryDBTypes, false, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}
	if err = randomize.Struct(seed, withdrawalHistoryTwo, withdrawalHistoryDBTypes, false, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = withdrawalHistoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = withdrawalHistoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WithdrawalHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func withdrawalHistoryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalHistory) error {
	*o = WithdrawalHistory{}
	return nil
}

func withdrawalHistoryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalHistory) error {
	*o = WithdrawalHistory{}
	return nil
}

func withdrawalHistoryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalHistory) error {
	*o = WithdrawalHistory{}
	return nil
}

func withdrawalHistoryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalHistory) error {
	*o = WithdrawalHistory{}
	return nil
}

func withdrawalHistoryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalHistory) error {
	*o = WithdrawalHistory{}
	return nil
}

func withdrawalHistoryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalHistory) error {
	*o = WithdrawalHistory{}
	return nil
}

func withdrawalHistoryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalHistory) error {
	*o = WithdrawalHistory{}
	return nil
}

func withdrawalHistoryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalHistory) error {
	*o = WithdrawalHistory{}
	return nil
}

func withdrawalHistoryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *WithdrawalHistory) error {
	*o = WithdrawalHistory{}
	return nil
}

func testWithdrawalHistoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &WithdrawalHistory{}
	o := &WithdrawalHistory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory object: %s", err)
	}

	AddWithdrawalHistoryHook(boil.BeforeInsertHook, withdrawalHistoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	withdrawalHistoryBeforeInsertHooks = []WithdrawalHistoryHook{}

	AddWithdrawalHistoryHook(boil.AfterInsertHook, withdrawalHistoryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	withdrawalHistoryAfterInsertHooks = []WithdrawalHistoryHook{}

	AddWithdrawalHistoryHook(boil.AfterSelectHook, withdrawalHistoryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	withdrawalHistoryAfterSelectHooks = []WithdrawalHistoryHook{}

	AddWithdrawalHistoryHook(boil.BeforeUpdateHook, withdrawalHistoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	withdrawalHistoryBeforeUpdateHooks = []WithdrawalHistoryHook{}

	AddWithdrawalHistoryHook(boil.AfterUpdateHook, withdrawalHistoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	withdrawalHistoryAfterUpdateHooks = []WithdrawalHistoryHook{}

	AddWithdrawalHistoryHook(boil.BeforeDeleteHook, withdrawalHistoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	withdrawalHistoryBeforeDeleteHooks = []WithdrawalHistoryHook{}

	AddWithdrawalHistoryHook(boil.AfterDeleteHook, withdrawalHistoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	withdrawalHistoryAfterDeleteHooks = []WithdrawalHistoryHook{}

	AddWithdrawalHistoryHook(boil.BeforeUpsertHook, withdrawalHistoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	withdrawalHistoryBeforeUpsertHooks = []WithdrawalHistoryHook{}

	AddWithdrawalHistoryHook(boil.AfterUpsertHook, withdrawalHistoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	withdrawalHistoryAfterUpsertHooks = []WithdrawalHistoryHook{}
}

func testWithdrawalHistoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WithdrawalHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWithdrawalHistoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(withdrawalHistoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := WithdrawalHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWithdrawalHistoryToManyWithdrawalCryptos(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WithdrawalHistory
	var b, c WithdrawalCrypto

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, withdrawalCryptoDBTypes, false, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, withdrawalCryptoDBTypes, false, withdrawalCryptoColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.WithdrawalHistoryID = a.ID
	c.WithdrawalHistoryID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.WithdrawalCryptos().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.WithdrawalHistoryID == b.WithdrawalHistoryID {
			bFound = true
		}
		if v.WithdrawalHistoryID == c.WithdrawalHistoryID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := WithdrawalHistorySlice{&a}
	if err = a.L.LoadWithdrawalCryptos(ctx, tx, false, (*[]*WithdrawalHistory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.WithdrawalCryptos); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.WithdrawalCryptos = nil
	if err = a.L.LoadWithdrawalCryptos(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.WithdrawalCryptos); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testWithdrawalHistoryToManyWithdrawalFiats(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WithdrawalHistory
	var b, c WithdrawalFiat

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, withdrawalFiatDBTypes, false, withdrawalFiatColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, withdrawalFiatDBTypes, false, withdrawalFiatColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.WithdrawalHistoryID = a.ID
	c.WithdrawalHistoryID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.WithdrawalFiats().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.WithdrawalHistoryID == b.WithdrawalHistoryID {
			bFound = true
		}
		if v.WithdrawalHistoryID == c.WithdrawalHistoryID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := WithdrawalHistorySlice{&a}
	if err = a.L.LoadWithdrawalFiats(ctx, tx, false, (*[]*WithdrawalHistory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.WithdrawalFiats); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.WithdrawalFiats = nil
	if err = a.L.LoadWithdrawalFiats(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.WithdrawalFiats); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testWithdrawalHistoryToManyAddOpWithdrawalCryptos(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WithdrawalHistory
	var b, c, d, e WithdrawalCrypto

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, withdrawalHistoryDBTypes, false, strmangle.SetComplement(withdrawalHistoryPrimaryKeyColumns, withdrawalHistoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*WithdrawalCrypto{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, withdrawalCryptoDBTypes, false, strmangle.SetComplement(withdrawalCryptoPrimaryKeyColumns, withdrawalCryptoColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*WithdrawalCrypto{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddWithdrawalCryptos(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.WithdrawalHistoryID {
			t.Error("foreign key was wrong value", a.ID, first.WithdrawalHistoryID)
		}
		if a.ID != second.WithdrawalHistoryID {
			t.Error("foreign key was wrong value", a.ID, second.WithdrawalHistoryID)
		}

		if first.R.WithdrawalHistory != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.WithdrawalHistory != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.WithdrawalCryptos[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.WithdrawalCryptos[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.WithdrawalCryptos().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testWithdrawalHistoryToManyAddOpWithdrawalFiats(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WithdrawalHistory
	var b, c, d, e WithdrawalFiat

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, withdrawalHistoryDBTypes, false, strmangle.SetComplement(withdrawalHistoryPrimaryKeyColumns, withdrawalHistoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*WithdrawalFiat{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, withdrawalFiatDBTypes, false, strmangle.SetComplement(withdrawalFiatPrimaryKeyColumns, withdrawalFiatColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*WithdrawalFiat{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddWithdrawalFiats(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.WithdrawalHistoryID {
			t.Error("foreign key was wrong value", a.ID, first.WithdrawalHistoryID)
		}
		if a.ID != second.WithdrawalHistoryID {
			t.Error("foreign key was wrong value", a.ID, second.WithdrawalHistoryID)
		}

		if first.R.WithdrawalHistory != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.WithdrawalHistory != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.WithdrawalFiats[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.WithdrawalFiats[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.WithdrawalFiats().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testWithdrawalHistoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWithdrawalHistoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WithdrawalHistorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWithdrawalHistoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := WithdrawalHistories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	withdrawalHistoryDBTypes = map[string]string{`ID`: `TEXT`, `Exchange`: `TEXT`, `ExchangeID`: `TEXT`, `Status`: `TEXT`, `Currency`: `TEXT`, `Amount`: `REAL`, `Description`: `TEXT`, `WithdrawType`: `INTEGER`, `CreatedAt`: `TIMESTAMP`, `UpdatedAt`: `TIMESTAMP`}
	_                        = bytes.MinRead
)

func testWithdrawalHistoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(withdrawalHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(withdrawalHistoryAllColumns) == len(withdrawalHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WithdrawalHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testWithdrawalHistoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(withdrawalHistoryAllColumns) == len(withdrawalHistoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &WithdrawalHistory{}
	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WithdrawalHistories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, withdrawalHistoryDBTypes, true, withdrawalHistoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WithdrawalHistory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(withdrawalHistoryAllColumns, withdrawalHistoryPrimaryKeyColumns) {
		fields = withdrawalHistoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			withdrawalHistoryAllColumns,
			withdrawalHistoryPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := WithdrawalHistorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
