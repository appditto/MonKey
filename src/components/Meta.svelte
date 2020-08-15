<script>
  export let metadata = {};
</script>

<svelte:head>
  {#if metadata.title}
    <title>{metadata.title}</title>
    <meta name="title" content={metadata.title} />
    <meta property="og:title" content={metadata.title} />
    <meta property="twitter:title" content={metadata.title} />
  {/if}

  {#if metadata.description}
    <meta name="description" content={metadata.description} />
    <meta property="og:description" content={metadata.description} />
    <meta property="twitter:description" content={metadata.description} />
  {/if}

  {#if metadata.image}
    <meta property="og:image" content={metadata.image} />
    <meta property="twitter:image" content={metadata.image} />
  {/if}

  {#if metadata.imageAlt}
    <meta property="og:image:alt" content={metadata.imageAlt} />
    <meta property="twitter:image:alt" content={metadata.imageAlt} />
  {/if}

  {#if metadata.url}
    <meta property="og:url" content={metadata.url} />
    <meta property="twitter:url" content={metadata.url} />
    <link rel="canonical" href={metadata.canonical || metadata.url} />
  {/if}

  {#if metadata.type}
    <meta property="og:type" content={'website' || metadata.type} />
  {/if}

  <meta
    property="twitter:card"
    content={metadata.twitterCard || 'summary_large_image'}
  />
</svelte:head>
