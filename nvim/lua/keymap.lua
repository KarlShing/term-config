-- auto fill source with coc
vim.api.nvim_set_keymap('i', '<CR>', 'coc#pum#visible() ? coc#pum#confirm() : "\\<CR>"', { noremap = true, expr = true })

-- go to definition with coc
vim.api.nvim_set_keymap('n', '<C-M>', '<Plug>(coc-definition)', { noremap = true, silent = true})
