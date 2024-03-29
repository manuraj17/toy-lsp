-- Connecting to an lsp via lua for nvim
vim.lsp.start({
  name = "toy-lsp",
  cmd = { "./bin/toy-lsp" },
  root_dir = vim.fn.getcwd(),
})
