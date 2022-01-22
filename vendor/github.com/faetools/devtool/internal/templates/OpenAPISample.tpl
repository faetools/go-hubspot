openapi: 3.0.0
info:
  description: '{{.Name}}'
  title: '{{.Name}}'
  version: '{{.Version}}'
servers:
  - url: https://{{.Name}}.fae-platform.com
    description: Production environment
  # - url: https://{{.Name}}.fae-staging.com
  #   description: Staging environment
  # - url: https://{{.Name}}-{namespace}.fae-dev.com
  #   description: Development environment
  #   variables:
  #     namespace:
  #       default: username
  #       description: The namespace of the development service, usually your username.
paths:
  /user/{id}:
    get:
      summary: Get a user by ID
      operationId: UserGet
      parameters:
        - name: id
          in: path
          description: User ID
          required: true
          schema:
            type: string
      responses:
        '200':
          $ref: '#/components/responses/UserResponse'
        '404':
          description: User not found
  /user:
    post:
      summary: Add a new user
      operationId: UserAdd
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/User'
      responses:
        '201':
          description: User successfully added
components:
  # securitySchemes:
  #   ApiKeyAuth:
  #     type: apiKey
  #     in: header
  #     name: X-Fae-Api-Key
  schemas:
    User:
      type: object
      required:
        - first_name
        - last_name
        - email
      properties:
        first_name:
          type: string
        last_name:
          type: string
        email:
          type: string
  responses:
        UserResponse:
            description: Returns a user.
            content:
                application/json:
                    schema:
                        $ref: '#/components/schemas/User'
# security:
#   - ApiKeyAuth: []
