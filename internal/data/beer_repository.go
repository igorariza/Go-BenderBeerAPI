package data

import (
	"context"
	"log"
	"time"

	beerv1alpha1 "github.com/igorariza/gogrpc-apibeer/gen/go"
)

type BeerRepository struct {
	Data *Data
}

func (pr *BeerRepository) Create(ctx context.Context, p *beerv1alpha1.CreateBeerRequest) (*beerv1alpha1.CreateBeerResponse, error) {
	q := `
    INSERT INTO beers (created_at, updated_at, name, brewery, country, price, currency)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id;
    `

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, time.Now(), time.Now(), "Balboa", "Brewery", "Country", 5, 1)
	if row != nil {
		log.Println("row: ", row)
	}
	return &beerv1alpha1.CreateBeerResponse{
		Msg: "Beer created",
	}, nil
}

func (pr *BeerRepository) List(ctx context.Context, p *beerv1alpha1.ListBeersRequest) (*beerv1alpha1.ListBeersResponse, error) {
	q := `
    INSERT INTO vaultv1alpha1 (created_at, updated_at, name)
        VALUES ($1, $2, $3)
        RETURNING id;
    `

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, time.Now(), time.Now(), p.Beer.Name)

	err = row.Scan(&p.Beer.Id)
	if err != nil {
		return nil, err
	}

	return &beerv1alpha1.ListBeersResponse{
		Msg: "Beer created",
	}, nil
}

func (pr *BeerRepository) Update(ctx context.Context, p *beerv1alpha1.UpdateBeerRequest) (*beerv1alpha1.UpdateBeerResponse, error) {
	q := `
    INSERT INTO vaultv1alpha1 (created_at, updated_at, name)
        VALUES ($1, $2, $3)
        RETURNING id;
    `

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, time.Now(), time.Now(), p.Beer.Name)

	err = row.Scan(&p.Beer.Id)
	if err != nil {
		return nil, err
	}

	return &beerv1alpha1.UpdateBeerResponse{
		Msg: "Beer created",
	}, nil
}

func (pr *BeerRepository) Get(ctx context.Context, p *beerv1alpha1.GetOneBeerRequest) (*beerv1alpha1.GetOneBeerResponse, error) {
	q := `
    INSERT INTO vaultv1alpha1 (created_at, updated_at, name)
        VALUES ($1, $2, $3)
        RETURNING id;
    `

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// row := stmt.QueryRowContext(ctx, time.Now(), time.Now(), p.Beer.Name)

	// err = row.Scan(&p.Beer.Id)
	// if err != nil {
	// 	return nil, err
	// }

	return &beerv1alpha1.GetOneBeerResponse{
		Msg: "Beer created",
	}, nil
}

func (pr *BeerRepository) Delete(ctx context.Context, p *beerv1alpha1.DeleteBeerRequest) (*beerv1alpha1.DeleteBeerResponse, error) {
	q := `
    INSERT INTO vaultv1alpha1 (created_at, updated_at, name)
        VALUES ($1, $2, $3)
        RETURNING id;
    `

	stmt, err := pr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// row := stmt.QueryRowContext(ctx, time.Now(), time.Now(), p.Beer.Name)

	// err = row.Scan(&p.Beer.Id)
	// if err != nil {
	// 	return nil, err
	// }

	return &beerv1alpha1.DeleteBeerResponse{
		Msg: "Beer created",
	}, nil
}
