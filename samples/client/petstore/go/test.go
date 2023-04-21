package main

import (
	"context"
	"fmt"

	sw "github.com/OpenAPITools/openapi-generator/samples/client/petstore/go/go-petstore"
	"github.com/go-openapi/swag"
)

func main() {
	ctx := context.Background()

	cfg := sw.NewConfiguration()
	cfg.Servers[0].URL = "http://127.0.0.1:8080/v2"

	s := sw.NewAPIClient(cfg)

	/**
	 * BODY VALUES
	 */

	// int32
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarInt32: swag.Int32(16)}).Execute(); resp.StatusCode == 201 {
		panic(fmt.Sprintf("should have failed because of minimum value 17: %d", resp.StatusCode))
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarInt32: swag.Int32(17)}).Execute(); resp.StatusCode != 407 {
		panic(fmt.Sprintf("should have succeed because of minimum value 17: %d", resp.StatusCode))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19: %v", err))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarInt32: swag.Int32(20)}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20: %v", err))
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarInt32: swag.Int32(42)}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42")
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarInt32: swag.Int32(43)}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42")
	}

	// []int32
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayInt32: []int32{16, 16}}).Execute(); resp.StatusCode == 201 {
		panic(fmt.Sprintf("should have failed because of minimum value 17: %d", resp.StatusCode))
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayInt32: []int32{17, 17}}).Execute(); resp.StatusCode != 407 {
		panic(fmt.Sprintf("should have succeed because of minimum value 17: %d", resp.StatusCode))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19: %v", err))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayInt32: []int32{20, 20}}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20: %v", err))
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayInt32: []int32{42, 42}}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42")
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayInt32: []int32{43, 43}}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42")
	}

	// int64
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarInt64: swag.Int64(16)}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of minimum value 17")
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarInt64: swag.Int64(17)}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of minimum value 17")
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19: %v", err))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarInt64: swag.Int64(20)}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20: %v", err))
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarInt64: swag.Int64(42)}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42")
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarInt64: swag.Int64(43)}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42")
	}

	// []int64
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayInt64: []int64{16, 16}}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of minimum value 17")
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayInt64: []int64{17, 17}}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of minimum value 17")
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19: %v", err))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayInt64: []int64{20, 20}}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20: %v", err))
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayInt64: []int64{42, 42}}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42")
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayInt64: []int64{43, 43}}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42")
	}

	// float32
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarFloat32: swag.Float32(17.16)}).Execute(); resp.StatusCode == 201 {
		panic(fmt.Sprintf("should have failed because of minimum value 17.17: %d", resp.StatusCode))
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarFloat32: swag.Float32(17.17)}).Execute(); resp.StatusCode != 407 {
		panic(fmt.Sprintf("should succeed failed because of minimum value 17.17: %d", resp.StatusCode))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19.19: %v", err))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarFloat32: swag.Float32(20.20)}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20.20: %v", err))
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarFloat32: swag.Float32(42.42)}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42.43")
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarFloat32: swag.Float32(42.43)}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42.43")
	}

	// []float32
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayFloat32: []float32{17.16, 17.16}}).Execute(); resp.StatusCode == 201 {
		panic(fmt.Sprintf("should have failed because of minimum value 17.17: %d", resp.StatusCode))
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayFloat32: []float32{17.17, 17.17}}).Execute(); resp.StatusCode != 407 {
		panic(fmt.Sprintf("should succeed failed because of minimum value 17.17: %d", resp.StatusCode))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19.19: %v", err))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayFloat32: []float32{20.20, 20.20}}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20.20: %v", err))
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayFloat32: []float32{42.42, 42.42}}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42.43")
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayFloat32: []float32{42.43, 42.43}}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42.43")
	}

	// float64
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarFloat64: swag.Float64(17.16)}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of minimum value 17.17")
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarFloat64: swag.Float64(17.17)}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of minimum value 17.17")
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19.19: %v", err))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarFloat64: swag.Float64(20.20)}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20.20: %v", err))
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarFloat64: swag.Float64(42.42)}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42.42")
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarFloat64: swag.Float64(42.43)}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42.42")
	}

	// []float64
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayFloat64: []float64{17.16, 17.16}}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of minimum value 17.17")
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayFloat64: []float64{17.17, 17.17}}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of minimum value 17.17")
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19.19: %v", err))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayFloat64: []float64{20.20, 20.20}}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20.20: %v", err))
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayFloat64: []float64{42.42, 42.42}}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42.42")
	}
	if resp, _ := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ArrayFloat64: []float64{42.43, 42.43}}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42.42")
	}

	// string
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be '19': %v", err))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarString: swag.String("20")}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be '20': %v", err))
	}
	if resp, err := s.PetApi.AddPet(ctx).Body(sw.Pet{PhotoUrls: []string{"x"}, Name: "x", ScalarString: swag.String("21")}).Execute(); resp.StatusCode != 407 {
		panic(fmt.Sprintf("passed value should be '21': %v", err))
	}

	// TODO: array

	/**
	 * PATH VALUES
	 */

	if _, resp, err := s.PetApi.GetPetById(ctx, 1, 20, 20, 20, 20).Execute(); resp.StatusCode != 501 {
		panic(fmt.Sprintf("should have failed succeed: %d / %+v", resp.StatusCode, err))
	}
	// Every other variation should fail because all values are out of bound
	if _, resp, _ := s.PetApi.GetPetById(ctx, 1, 16, 20, 20, 20).Execute(); resp.StatusCode == 501 {
		panic("should have failed")
	}
	if _, resp, _ := s.PetApi.GetPetById(ctx, 1, 43, 20, 20, 20).Execute(); resp.StatusCode == 501 {
		panic("should have failed")
	}
	if _, resp, _ := s.PetApi.GetPetById(ctx, 1, 20, 16, 20, 20).Execute(); resp.StatusCode == 501 {
		panic("should have failed")
	}
	if _, resp, _ := s.PetApi.GetPetById(ctx, 1, 20, 43, 20, 20).Execute(); resp.StatusCode == 501 {
		panic("should have failed")
	}
	if _, resp, _ := s.PetApi.GetPetById(ctx, 1, 20, 20, 17.16, 20).Execute(); resp.StatusCode == 501 {
		panic("should have failed")
	}
	if _, resp, _ := s.PetApi.GetPetById(ctx, 1, 20, 20, 42.43, 20).Execute(); resp.StatusCode == 501 {
		panic("should have failed")
	}
	if _, resp, _ := s.PetApi.GetPetById(ctx, 1, 20, 20, 20, 17.16).Execute(); resp.StatusCode == 501 {
		panic("should have failed")
	}
	if _, resp, _ := s.PetApi.GetPetById(ctx, 1, 20, 20, 20, 42.43).Execute(); resp.StatusCode == 501 {
		panic("should have failed")
	}

	/**
	 * QUERY VALUES
	 */

	// int32
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarInt32(16).Execute(); resp.StatusCode == 201 {
		panic(fmt.Sprintf("should have failed because of minimum value 17: %d", resp.StatusCode))
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarInt32(17).Execute(); resp.StatusCode != 407 {
		panic(fmt.Sprintf("should have succeed because of minimum value 17: %d", resp.StatusCode))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19: %v", err))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarInt32(20).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20: %v", err))
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarInt32(42).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42")
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarInt32(43).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42")
	}

	// []int32
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayInt32([]int32{16, 16}).Execute(); resp.StatusCode == 201 {
		panic(fmt.Sprintf("should have failed because of minimum value 17: %d", resp.StatusCode))
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayInt32([]int32{17, 17}).Execute(); resp.StatusCode != 407 {
		panic(fmt.Sprintf("should have succeed because of minimum value 17: %d", resp.StatusCode))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19: %v", err))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayInt32([]int32{20, 20}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20: %v", err))
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayInt32([]int32{42, 42}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42")
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayInt32([]int32{43, 43}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42")
	}

	// int64
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarInt64(16).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of minimum value 17")
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarInt64(17).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of minimum value 17")
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19: %v", err))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarInt64(20).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20: %v", err))
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarInt64(42).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42")
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarInt64(43).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42")
	}

	// []int64
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayInt64([]int64{16, 16}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of minimum value 17")
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayInt64([]int64{17, 17}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of minimum value 17")
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19: %v", err))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayInt64([]int64{20, 20}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20: %v", err))
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayInt64([]int64{42, 42}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42")
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayInt64([]int64{43, 43}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42")
	}

	// float32
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarFloat32(17.16).Execute(); resp.StatusCode == 201 {
		panic(fmt.Sprintf("should have failed because of minimum value 17.17: %d", resp.StatusCode))
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarFloat32(17.17).Execute(); resp.StatusCode != 407 {
		panic(fmt.Sprintf("should succeed failed because of minimum value 17.17: %d", resp.StatusCode))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19.19: %v", err))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarFloat32(20.20).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20.20: %v", err))
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarFloat32(42.42).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42.43")
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarFloat32(42.43).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42.43")
	}

	// []float32
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayFloat32([]float32{17.16, 17.16}).Execute(); resp.StatusCode == 201 {
		panic(fmt.Sprintf("should have failed because of minimum value 17.17: %d", resp.StatusCode))
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayFloat32([]float32{17.17, 17.17}).Execute(); resp.StatusCode != 407 {
		panic(fmt.Sprintf("should succeed failed because of minimum value 17.17: %d", resp.StatusCode))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19.19: %v", err))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayFloat32([]float32{20.20, 20.20}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20.20: %v", err))
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayFloat32([]float32{42.42, 42.42}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42.43")
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayFloat32([]float32{42.43, 42.43}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42.43")
	}

	// float64
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarFloat64(17.16).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of minimum value 17.17")
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarFloat64(17.17).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of minimum value 17.17")
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19.19: %v", err))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarFloat64(20.20).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20.20: %v", err))
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarFloat64(42.42).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42.42")
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarFloat64(42.43).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42.42")
	}

	// []float64
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayFloat64([]float64{17.16, 17.16}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of minimum value 17.17")
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayFloat64([]float64{17.17, 17.17}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of minimum value 17.17")
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be 19.19: %v", err))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayFloat64([]float64{20.20, 20.20}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be 20.20: %v", err))
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayFloat64([]float64{42.42, 42.42}).Execute(); resp.StatusCode != 407 {
		panic("should have succeed because of maximum value 42.42")
	}
	if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ArrayFloat64([]float64{42.43, 42.43}).Execute(); resp.StatusCode == 201 {
		panic("should have failed because of maximum value 42.42")
	}

	// string
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("default value should be '19': %v", err))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarString("20").Execute(); resp.StatusCode != 201 {
		panic(fmt.Sprintf("passed value should be '20': %v", err))
	}
	if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).ScalarString("21").Execute(); resp.StatusCode != 407 {
		panic(fmt.Sprintf("passed value should be '21': %v", err))
	}

	//if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"available"}).Execute(); resp.StatusCode != 501 {
	//	panic(fmt.Sprintf("should have failed succeed: %d / %+v", resp.StatusCode, err))
	//}

	// TODO array

	//
	//// test query param
	//if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"16"}).Execute(); resp.StatusCode == 501 {
	//	panic("should have failed")
	//}
	//if _, resp, _ := s.PetApi.FindPetsByStatus(ctx).Status([]string{"43"}).Execute(); resp.StatusCode == 501 {
	//	panic("should have failed")
	//}
	//if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"17"}).Execute(); resp.StatusCode != 501 {
	//	panic(err)
	//}
	//if _, resp, err := s.PetApi.FindPetsByStatus(ctx).Status([]string{"42"}).Execute(); resp.StatusCode != 501 {
	//	panic(err)
	//}

	// fmt.Println(s.StoreApi.GetOrderById(ctx, 2).Execute())
	// fmt.Println(s.StoreApi.GetOrderById(ctx, 0).Execute())
	// fmt.Println(s.StoreApi.GetOrderById(ctx, 20).Execute())
	//
	// // test POST(form)
	// s.PetApi.UpdatePetWithForm(ctx, 12830).Name("golang").Status("available")
	//
	// // test GET
	// resp, apiResponse, err := s.PetApi.GetPetById(ctx, 12830).Execute()
	// fmt.Println("GetPetById:", resp, err, apiResponse)
	//
	// err2, apiResponse2 := s.PetApi.DeletePet(ctx, 12830).Execute()
	// fmt.Println("DeletePet:", err2, apiResponse2)
}
