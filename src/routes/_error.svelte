<script>
  export let status;
  export let error;
  const dev = process.env.NODE_ENV === "development";
</script>

<svelte:head>
  <title>{status}</title>
</svelte:head>

<div class="w-full flex flex-col items-center justify-center my-auto pb-12 md:pb-32 px-4">
  <div class="y-container flex flex-row justify-center">
    <img
      class="w-full max-w-2xl h-auto"
      src="images/illustrations-foreground/error.svg"
      alt="Error Illustration"
    />
  </div>
  <p class="text-2xl px-4 text-center -mt-5 md:-mt-10">
    {status == '404' ? "The page you're looking for isn't there." : 'An error occurred.'}
  </p>
  <a
    href="/"
    class="{'bg-primary btn-primary border-black'} text-center w-64 max-w-full text-white text-xl font-bold
    rounded-xl border-2 px-3 md:px-6 py-2 mx-auto mt-5 md:mt-8"
  >
    Go Back Home
  </a>
  {#if dev && error.stack}
    <pre class="mt-6">{error.stack}</pre>
  {/if}
</div>
