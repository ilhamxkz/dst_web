import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import About from '../views/About.vue';
import Contact from '../views/Contact.vue';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
    meta: {
      title: 'CoreTech - IT Transformation & Modernization',
    },
  },
  {
    path: '/about',
    name: 'About',
    component: About,
    meta: {
      title: 'Page - CoreTech',
    },
  },
  {
    path: '/our-team',
    name: 'OurTeam',
    component: () => import('../views/OurTeam.vue'),
    meta: {
      title: 'Our Team - CoreTech',
    },
  },
  {
    path: '/case-study',
    name: 'CaseStudy',
    component: () => import('../views/CaseStudy.vue'),
    meta: {
      title: 'Case Study - CoreTech',
    },
  },
  {
    path: '/careers',
    name: 'Careers',
    component: () => import('../views/Careers.vue'),
    meta: {
      title: 'Careers - CoreTech',
    },
  },
  {
    path: '/services',
    name: 'Services',
    component: () => import('../views/Services.vue'),
    meta: {
      title: 'Our Services - CoreTech',
    },
  },
  {
    path: '/services/:id',
    name: 'ServiceDetail',
    component: () => import('@/views/ServiceDetail.vue'),
    meta: {
      title: 'Service Detail - CoreTech',
    },
  },

  {
    path: '/industry',
    name: 'Industry',
    component: () => import('../views/Industry.vue'),
    meta: {
      title: 'Industries - CoreTech',
    },
  },
  {
    path: '/industry/:id',
    name: 'IndustryDetail',
    component: () => import('../views/Industry.vue'),
    meta: {
      title: 'Industry Detail - CoreTech',
    },
  },
  {
    path: '/blog',
    name: 'Blog',
    component: () => import('../views/Blog.vue'),
    meta: {
      title: 'Blog - CoreTech',
    },
  },
  {
    path: '/pelatihan/:id',
    name: 'PelatihanDetail',
    component: () => import('../views/PelatihanDetail.vue'),
    meta: {
      title: 'Pelatihan Detail - CoreTech',
    },
  },
  {
    path: '/contact',
    name: 'Contact',
    component: Contact,
    meta: {
      title: 'Contact Us - CoreTech',
    },
  },
  // FAQ route removed - now using anchor link to FAQ section in Home page
  // {
  //   path: '/faq_home',
  //   name: 'FAQ',
  //   component: () => import('../views/FAQ.vue'),
  //   meta: {
  //     title: 'FAQ - CoreTech'
  //   }
  // },
  {
    path: '/404',
    name: '404Page',
    component: () => import('../views/NotFound.vue'),
    meta: {
      title: '404 - Page Not Found - CoreTech',
    },
  },
  // Catch-all route for 404
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('../views/NotFound.vue'),
    meta: {
      title: 'Page Not Found - CoreTech',
    },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) return savedPosition;

    // Handle hash anchors like #team, #faq on Home
    if (to.hash) {
      return {
        el: to.hash,
        behavior: 'smooth',
      };
    }

    // Default scroll to top
    return { top: 0 };
  },
});

// Update document title based on route meta
router.beforeEach((to, from, next) => {
  if (to.meta && to.meta.title) {
    document.title = to.meta.title;
  }
  next();
});

export default router;
