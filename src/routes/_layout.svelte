<script>
  import Navbar from "../components/Navbar.svelte";
  import Footer from "../components/Footer.svelte";
  import Tailwindcss from "../components/Tailwindcss.svelte";
  import GoogleAnalytics from "../components/GoogleAnalytics.svelte";
  import { stores } from "@sapper/app";

  export let segment;

  let trackingId = "UA-145357881-6";
</script>

<svelte:head>
  <link rel="apple-touch-icon" href="/mark-ios-192.png" />
</svelte:head>

<GoogleAnalytics {stores} {trackingId} />
<Tailwindcss />
<div class="min-h-screen flex flex-col justify-between">
  <Navbar {segment} />
  <main class="my-auto relative z-10">
    <slot />
  </main>
  <Footer />
</div>
