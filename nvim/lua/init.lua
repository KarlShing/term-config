require('plugins')

vim.api.nvim_create_autocmd("BufWritePre", {
  pattern = "*.go",
  callback = function()
    require('go.format').gofmt()
  end,
  group = format_sync_grp
})

require('keymap')

require('go').setup()
