package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Item struct { // Создаем структуру товаров
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Categories string `json:"categories"`
}

var items = []Item{} // Создаем глобальную переменную для хранения товаров
var nextID = 1       // Счетчик для увелечения при добавлении новых товаров

func main() {
	app := fiber.New(fiber.Config{ // Создаем новое приложение на Fiber
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(500).JSON(fiber.Map{"error": "Внутренняя ошибка сервера. Ошибка 500 - Internal Server Error"})
		},
	})

	// POST /item - создание товара
	app.Post("/item", func(c *fiber.Ctx) error { // Обработчик для POST-запросов
		var newItem Item

		if err := c.BodyParser(&newItem); err != nil { // Чтение и проверка на правильность JSON
			return c.Status(400).JSON(fiber.Map{"error": "Неверный JSON. Ошибка 400 - Bad Request"})
		}

		if newItem.Name == "" || newItem.Price <= 0 { // Проверка на заполнения имени и цены
			return c.Status(400).JSON(fiber.Map{"error": "Имя и цена обязательны. Ошибка 400 - Bad Request"})
		}

		newItem.ID = nextID            // Присваиваем уникальный ID
		items = append(items, newItem) // Добавляем товар в глобальную переменную (список)
		nextID++                       // Увеличиваем счетчик ID

		return c.Status(200).JSON(fiber.Map{"id": newItem.ID}) // Код 200 - Все хорошо
	})

	// GET /item/:id - получить товар по ID
	app.Get("/item/:id", func(c *fiber.Ctx) error { // Обработчик для GET-запросов
		idParam := c.Params("id") // Получение ID из URL

		id, err := strconv.Atoi(idParam) // Преобразование строки в число (string > int)
		if err != nil {                  // Проверка на ошибку при преобразовании строки в число
			return c.Status(400).JSON(fiber.Map{"error": "ID должен быть числом. Ошибка 400 - Bad Request"})
		}

		for _, item := range items { // Перебираем все товары в списке, для нахождения нужного ID ("_" - индекс; "range"- перебор элементов)
			if item.ID == id {
				return c.JSON(item) // Нашли товар - возращаем
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Товар не найден. Ошибка 404 - Not Found"}) // Проверка на ошибку, если товар не найден
	})

	// PUT /item/:id - обновить товар по ID
	app.Put("/item/:id", func(c *fiber.Ctx) error { // Обработчик для PUT-запроса
		idParam := c.Params("id") // Получаем ID из URL

		id, err := strconv.Atoi(idParam) // Преобразование строки в число
		if err != nil {                  // Проверка на преобразование строки в число
			return c.Status(400).JSON(fiber.Map{"error": "ID должен быть числом. Ошибка 400 - Bad Request"})
		}

		for i, item := range items { // Ищем товар с таким ID
			if item.ID == id {
				var updatedItem Item // Создаем переменную для новых данных

				if err := c.BodyParser(&updatedItem); err != nil { // Чтение и проверка на правильность JSON
					return c.Status(400).JSON(fiber.Map{"error": "Неверный JSON. Ошибка 400 - Bad Request"})
				}

				if updatedItem.Name == "" || updatedItem.Price <= 0 { // Проверка на заполнение имени и цены
					return c.Status(400).JSON(fiber.Map{"error": "Имя и цена обязательны. Ошибка 400 - Bad Request"})
				}

				updatedItem.ID = id                // Сохраняем ID
				items[i] = updatedItem             // Обновляем товар
				return c.JSON(fiber.Map{"id": id}) // Возращаем ID
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Товар не найден. Ошибка 404 - Not Found"}) // Проверка на ошибку, если товар не найден
	})

	// DELETE /item/:id - удаление товара по ID
	app.Delete("/item/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id") // Получаем ID из URL

		id, err := strconv.Atoi(idParam) // Преобразование строки в число
		if err != nil {                  // Проверка на преобразование строки в число
			return c.Status(400).JSON(fiber.Map{"error": "ID должен быть числом. Ошибка 400 - Bad Request"})
		}

		for i, item := range items { // Ищем товар с таким ID
			if item.ID == id {
				items = append(items[:i], items[i+1:]...) // Удаляем товар из списка
				// items[:i] — всё до найденного товара (не включая его)
				// items[i+1:] — всё после найденного товара
				// ... — распаковка: объединяем два куска в один слайс
				return c.JSON(fiber.Map{"id": id}) // Возращаем успех
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Товар не найден. Ошибка 404 - Not Found"})
	})
	// app.Listen(":3000")
}
