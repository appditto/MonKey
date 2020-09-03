<script>
  export let stores;
  export let trackingId;

  if (typeof window !== "undefined") {
    window.dataLayer = window.dataLayer || [];
    window.gtag = function gtag() {
      window.dataLayer.push(arguments);
    };
    window.gtag("js", new Date());
    window.gtag("config", trackingId, { send_page_view: false });
  }
  const { page } = stores();
  $: {
    if (typeof gtag !== "undefined") {
      window.gtag("config", trackingId, {
        page_path: $page.path,
      });
    }
  }
</script>

<svelte:head>
  <script async src="https://www.googletagmanager.com/gtag/js?id={trackingId}">

  </script>
</svelte:head>
