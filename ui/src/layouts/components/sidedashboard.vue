<script setup>
import VerticalNavLink from '@layouts/components/VerticalNavLink.vue'
</script>
<template >
    <VerticalNavLink
    :item="{
      title: 'Dashbord',
      icon: 'ri-dashboard-line',
      to: '/dashboard',
    }"
  />
</template>