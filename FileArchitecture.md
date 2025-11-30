Файловая архитектура
```
.
├── cmd
│   └── server
│       └── main.go
├── config
│   └── local.yaml
├── docker-compose.yaml
├── Dockerfile
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── example.env
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── app.go
│   │   └── migrations
│   │       ├── 20251128203423_init_schema.sql
│   │       ├── 20251129003923_add_test_session.sql
│   │       └── 20251129221957_update_test_session.sql
│   ├── delivery
│   │   └── http_server
│   │       ├── router.go
│   │       └── v1
│   │           ├── dto.go
│   │           ├── handler.go
│   │           ├── question.go
│   │           ├── question_test.go
│   │           ├── session.go
│   │           └── template.go
│   ├── domain
│   │   ├── auth_roles.go
│   │   ├── context_helpers.go
│   │   ├── domain.go
│   │   ├── question.go
│   │   ├── question_test.go
│   │   ├── test_session.go
│   │   ├── test_session_test.go
│   │   ├── test_template.go
│   │   ├── test_template_test.go
│   │   ├── types.go
│   │   └── types_test.go
│   ├── mocks
│   │   ├── QuestionRepository.go
│   │   ├── QuestionUseCase.go
│   │   ├── SessionUseCase.go
│   │   ├── TemplateRepository.go
│   │   ├── TemplateUseCase.go
│   │   └── TestSessionRepository.go
│   ├── parseconfig
│   │   └── config.go
│   ├── repository
│   │   ├── memory
│   │   │   ├── memory.go
│   │   │   ├── questions_repository.go
│   │   │   └── template.go
│   │   └── postgres
│   │       ├── question.go
│   │       ├── question_test.go
│   │       ├── session.go
│   │       ├── template.go
│   │       └── testing_utils_test.go
│   └── usecase
│       ├── ports
│       │   ├── mocks_test.go
│       │   ├── question_repository.go
│       │   ├── session_repository.go
│       │   └── template_repository.go
│       ├── question.go
│       ├── question_test.go
│       ├── session.go
│       ├── template.go
│       └── template_test.go
├── makefile
├── pkg
│   ├── logger
│   │   └── logger.go
│   └── postgres
│       └── connection.go
├── README.md
└── TASK.md
```