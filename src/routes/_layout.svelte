<script>
  import Navbar from "../components/Navbar.svelte";
  import Footer from "../components/Footer.svelte";
  import Tailwindcss from "../components/Tailwindcss.svelte";
  export let segment;
</script>

<svelte:head>
  <link rel="apple-touch-icon" href="/mark-ios-192.png" />
</svelte:head>

<Tailwindcss />
<div class="min-h-screen flex flex-col justify-between">
  <Navbar {segment} />
  <main class="my-auto relative z-10">
    <slot />
  </main>
  <Footer />
</div>
