<script setup>
import { RouterLink, RouterView } from 'vue-router';
</script>

<template>
  <header>
    <div class="nav-wrapper">
      <nav>
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/code">Code</RouterLink>
        <RouterLink to="/flow">Flow</RouterLink>
      </nav>
    </div>
  </header>
  <main>
    <RouterView />
  </main>
</template>

<style>
@import '@/assets/base.css';

#app {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
}

header {
  line-height: 3rem;
}

.nav-wrapper {
  display: flex;
  justify-content: start;
  align-items: flex-start;
  background-color: #ebf0f1;
}

main {
  height: calc(100% - 3rem);
}

a,
.green {
  text-decoration: none;
  color: hsla(160, 100%, 37%, 1);
  transition: 0.4s;
}

@media (hover: hover) {
  a:hover {
    background-color: hsla(160, 100%, 37%, 0.2);
  }
}

nav {
  width: 100%;
  font-size: 14px;
  text-align: center;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}
</style>
