package main

import "github.com/gofiber/fiber/v2"

// @Summary      Lista todos os produtos
// @Description  Retorna uma lista de produtos
// @Tags         products
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {array} Product
// @Router       /api/products [get]
func getProducts(c *fiber.Ctx) error {
	var products []Product
	DB.Find(&products)
	return c.JSON(products)
}

// @Summary      Obtém um produto
// @Description  Retorna os detalhes de um produto pelo ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path      int  true  "ID do Produto"
// @Success      200  {object}  Product
// @Failure      404  {object}  map[string]string
// @Router       /api/products/{id} [get]
func getProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product Product
	if err := DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(product)
}

// @Summary      Cria um produto
// @Description  Cria um novo produto
// @Tags         products
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        product body Product true "Dados do Produto"
// @Success      201 {object} Product
// @Failure      400 {object} map[string]string
// @Router       /api/products [post]
func createProduct(c *fiber.Ctx) error {
	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	DB.Create(&product)
	return c.JSON(product)
}

// @Summary      Atualiza um produto
// @Description  Atualiza os dados de um produto existente
// @Tags         products
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id      path   int     true  "ID do Produto"
// @Param        product body   Product true  "Dados Atualizados"
// @Success      200 {object} Product
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /api/products/{id} [put]
func updateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product Product
	if err := DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	DB.Save(&product)
	return c.JSON(product)
}

// @Summary      Deleta um produto
// @Description  Remove um produto pelo ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id   path   int  true  "ID do Produto"
// @Success      204  {object}  nil
// @Failure      500  {object}  map[string]string
// @Router       /api/products/{id} [delete]
func deleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := DB.Delete(&Product{}, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.SendStatus(204)
}
