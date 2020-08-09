<script>
  import Navbar from "../components/Navbar.svelte";
  import Footer from "../components/Footer.svelte";
  export let segment;
</script>

<div class="min-h-screen flex flex-col justify-between">
  <Navbar {segment} />
  <main class="my-auto relative z-10">
    <slot />
  </main>
  <Footer />
</div>
