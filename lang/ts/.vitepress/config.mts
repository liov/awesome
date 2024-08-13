import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "awesome",
  description: "语言学习,环境配置,开发笔记",
  srcDir: '../..',
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'note', link: '/note/' },
      { text: 'repost', link: '/repost/' }
    ],

    sidebar: [
      {
        text: 'note',
        base: '/note/',
        items: [
          { text: 'ASCII', link: '/ASCII' },
          { text: 'IEEE754', link: '/IEEE754' },
          { text: '我遇到过的那些坑', link: '/我遇到过的那些坑' },
        ]
      },
      {
        text: 'repost',
        base: '/repost/',
        items: [
          { text: '关于DDD', link: '/关于DDD' },
          { text: '补充DDD', link: '/补充DDD' }
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/liov/awesome' }
    ]
  }
})
