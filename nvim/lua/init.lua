require('plugins')

vim.api.nvim_create_user_command("IM", "call CocAction('runCommand', 'editor.action.organizeImport')", {})
vim.api.nvim_create_user_command("FMT", "call CocAction('format')", {})
vim.api.nvim_create_autocmd("BufWritePre", {
  pattern = "*.go",
  callback = function()
    vim.cmd("IM")
    vim.cmd("FMT")
  end,
  group = format_sync_grp
})

require('keymap')

require('go').setup()
