package api

// setKey godoc
// @Summary Set key-value pair
// @Description set key-value pair with optional TTL
// @Tags key-value store
// @Accept json
// @Produce json
// @Param request body SetValueRequest true "Request body"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /set [post]
func setKeyDoc() {}

// getKey godoc
// @Summary Get value by key
// @Description get value associated with the key
// @Tags key-value store
// @Produce  json
// @Param   key     path      string     true  "Key"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /get/{key} [get]
func getKeyDoc() {}

// deleteKey godoc
// @Summary Delete key-value pair
// @Description delete key-value pair by key
// @Tags key-value store
// @Produce  json
// @Param   key     path      string     true  "Key"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /delete/{key} [delete]
func deleteKeyDoc() {}
