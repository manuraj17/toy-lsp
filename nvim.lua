vim.lsp.start({
  name = "emoji-lsp",
  cmd = { "./bin/toy-lsp" },
  root_dir = vim.fn.getcwd(),
})
