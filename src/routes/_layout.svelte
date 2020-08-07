<script>
  import Nav from "../components/Nav.svelte";
  import Navbar from "../components/Navbar.svelte";
  export let segment;
</script>

<Navbar {segment} />

<main>
  <slot />
</main>
