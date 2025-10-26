export const routes = [
  // Redirect root ke login dulu
  { path: '/', redirect: '/login' },

  // Layout utama (setelah login)
  {
    path: '/',
    component: () => import('@/layouts/default.vue'),
    children: [
      {
        path: 'comment',
        component: () => import('@/pages/blog_comments.vue'),
      },
      {
        path: 'faq',
        component: () => import('@/pages/faq.vue'),
      },
      {
        path: 'post',
        component: () => import('@/pages/blog_post.vue'),
      },
      {
        path: 'menu',
        component: () => import('@/pages/menu.vue'),
      },
      {
        path: 'editor',
        component: () => import('@/pages/editor.vue'),
      },
      {
        path: 'dashboard',
        component: () => import('@/pages/dashboard.vue'),
      },
      {
        path: 'account-settings',
        component: () => import('@/pages/Role.vue'),
      },
      {
        path: 'akses',
        component: () => import('@/pages/akses.vue'),
      },
      {
        path: 'icons',
        component: () => import('@/pages/icons.vue'),
      },
      {
        path: 'cards',
        component: () => import('@/pages/cards.vue'),
      },
      {
        path: 'tables',
        component: () => import('@/pages/tables.vue'),
      },
      {
        path: 'form-layouts',
        component: () => import('@/pages/form-layouts.vue'),
      },
      {
        path: 'contact',
        component: () => import('@/pages/contact.vue'),
      },
      {
        path: 'companyprofile',
        component: () => import('@/pages/companyprofile.vue'),
      },
      {
        path: 'message',
        component: () => import('@/pages/message.vue'),
      },
      {
        path: 'blog-posts',
        component: () => import('@/pages/blog_post.vue'),
      },
      {
        path: 'blog-comments',
        component: () => import('@/pages/blog_comments.vue'),
      },
      {
        path: 'roles',
        component: () => import('@/pages/Role.vue'),
      },
      {
        path: 'clients',
        component: () => import('@/pages/clients.vue'),
      },
      {
        path: 'teams',
        component: () => import('@/pages/team.vue'),
      },
      {
        path: 'services',
        component: () => import('@/pages/services.vue'),
      },
      {
        path: 'project_images',
        component: () => import('@/pages/project-images.vue'),
      },
      {
        path: 'projects',
        component: () => import('@/pages/projects.vue'),
      },
      {
        path: 'experiences',
        component: () => import('@/pages/experiences.vue'),
      },
      {
        path: 'testimoni',
        component: () => import('@/pages/testimoni.vue'),
      },
      {
        path: 'hero_background',
        component: () => import('@/pages/hero_background.vue'),
      },
      {
        path: 'pelatihan',
        component: () => import('@/pages/pelatihan.vue'),
      }

    ],
  },

  // Layout blank (login, register, error page)
  {
    path: '/',
    component: () => import('@/layouts/blank.vue'),
    children: [
      {
        path: 'login',
        component: () => import('@/pages/login.vue'),
      },
      {
        path: 'register',
        component: () => import('@/pages/register.vue'),
      },
      {
        path: '/:pathMatch(.*)*',
        component: () => import('@/pages/[...error].vue'),
      },
    ],
  },
]
