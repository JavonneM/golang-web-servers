local dap = require("dap")
dap.configurations.go = {
     {
        name = "Debug",
        type = "go",
        request = "launch",
        program = "./cmd/myapp/main.go",
        cwd = "${workspaceFolder}",
        stopOnEntry = false,
    },
    {
            type = "go",
            name = "Debug Test",
            request = "launch",
            mode = "test",
            program = "internal/myapp/...",
            args = function()
                local argument_string = vim.fn.input('Program arguments: ')
                return {"-v", "-run ", argument_string}
            end
        }
}

