vim.cmd [[packadd packer.nvim]]

return require('packer').startup(function(use)
  -- self 
  use 'wbthomason/packer.nvim'
  
  -- LSP configs
  use 'neovim/nvim-lspconfig'
  use {
    'neoclide/coc.nvim',
	branch = 'release'
  }

  -- debug
  use 'mfussenegger/nvim-dap'
  use { "rcarriga/nvim-dap-ui", requires = {"mfussenegger/nvim-dap"} }

  -- style
  use 'nvim-treesitter/nvim-treesitter'

  -- tools
  use {
    "windwp/nvim-autopairs",
    config = function() require("nvim-autopairs").setup {} end
  }

  -- golang
  use 'ray-x/go.nvim'
end)
