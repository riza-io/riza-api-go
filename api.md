# Secrets

Response Types:

- <a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#Secret">Secret</a>
- <a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#SecretListResponse">SecretListResponse</a>

Methods:

- <code title="post /v1/secrets">client.Secrets.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#SecretService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#SecretNewParams">SecretNewParams</a>) (<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#Secret">Secret</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/secrets">client.Secrets.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#SecretService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#SecretListResponse">SecretListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tools

Response Types:

- <a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#Tool">Tool</a>
- <a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#ToolListResponse">ToolListResponse</a>
- <a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#ToolExecResponse">ToolExecResponse</a>

Methods:

- <code title="post /v1/tools">client.Tools.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#ToolService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#ToolNewParams">ToolNewParams</a>) (<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#Tool">Tool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/tools/{id}">client.Tools.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#ToolService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#ToolUpdateParams">ToolUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#Tool">Tool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tools">client.Tools.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#ToolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#ToolListResponse">ToolListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/tools/{id}/execute">client.Tools.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#ToolService.Exec">Exec</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#ToolExecParams">ToolExecParams</a>) (<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#ToolExecResponse">ToolExecResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tools/{id}">client.Tools.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#ToolService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#Tool">Tool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Command

Response Types:

- <a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#CommandExecResponse">CommandExecResponse</a>

Methods:

- <code title="post /v1/execute">client.Command.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#CommandService.Exec">Exec</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#CommandExecParams">CommandExecParams</a>) (<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go">riza</a>.<a href="https://pkg.go.dev/github.com/riza-io/riza-api-go#CommandExecResponse">CommandExecResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
