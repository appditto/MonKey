<script>
  import Navbar from "../components/Navbar.svelte";
  import Footer from "../components/Footer.svelte";
  export let segment;
</script>

<svelte:head>
  <link rel="apple-touch-icon" href="/mark-ios-192.png" />
</svelte:head>

<div class="min-h-screen flex flex-col justify-between">
  <Navbar {segment} />
  <main class="my-auto relative z-10">
    <slot />
  </main>
  <Footer />
</div>
