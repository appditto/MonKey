<script>
  import LazyImage from "../components/LazyImage.svelte";
  import axios from "axios";
  import { onMount } from "svelte";
  import Meta from "../components/Meta.svelte";
  const metadata = {
    title: "MonKey | Documentation",
    description: "MonKey API Documentation",
    image: "https://monkey.banano.cc/previews/documentation-preview.png",
    imageAlt: "Address visualisation for Banano",
    url: "https://monkey.banano.cc/documentation",
  };
  let callSvgText;
  async function getSvgText() {
    try {
      return axios.get("/monkeyText.json");
    } catch (e) {
      console.error(e);
    }
  }
  onMount(async () => {
    const res = await getSvgText();
    callSvgText = res.data.data;
  });
</script>

<!-- HTML -->
<Meta {metadata} />
<div class="y-container flex flex-col items-center px-4 pt-20">
  <h1
    class="font-bold text-4xl md:text-6xl mx-5 md:mx-12 text-center leading-tight tracking-tight
    mt-5 md:mt-16 lg:mt-20"
  >
    API Documentation
  </h1>
  <p class="text-center mx-5 mt-4 text-2xl">
    We have a single endpoint and just a couple of paramaters.
  </p>
  <div class="divider mt-20 mb-16" />
  <!-- Endpoint -->
  <h2 class="font-bold text-4xl md:text-5xl mx-4 text-center leading-tight tracking-tight">
    Endpoint
  </h2>
  <div class="flex flex-wrap justify-center mt-8">
    <code class="text-3xl bg-orangeLight rounded-lg px-3 py-1 font-bold m-2">get</code>
    <div class="flex items-center m-2 text-orangeLight bg-black px-4 py-2 rounded-lg mono">
      <span class="text-0 break-all">
        <span class="text-xl">https://monkey.banano.cc/api/v1/</span>
        <span class="text-xl text-purpleLight">{'<address>'}</span>
      </span>
    </div>
  </div>
  <div class="divider mt-20 mb-16" />
  <!-- Example Call -->
  <h3 class="font-medium text-3xl md:text-4xl mx-4 text-center leading-tight tracking-tight">
    Example Call
  </h3>
  <div class="flex flex-wrap justify-center items-center mt-5">
    <code class="text-3xl bg-orangeLight rounded-lg px-3 py-1 font-bold m-2">get</code>
    <div
      class="w-full max-w-2xl flex items-center m-2 text-orangeLight bg-black px-4 py-2 rounded-lg
      mono"
    >
      <span class="text-0 break-all">
        <span class="text-xl">https://monkey.banano.cc/api/v1/</span>
        <span class="text-xl text-purpleLight">
          ban_1ka1ium4pfue3uxtntqsrib8mumxgazsjf58gidh1xeo5te3whsq8z476goo
        </span>
      </span>
    </div>
  </div>
  <!-- Arrow Down -->
  <div class="w-12 h-12 my-4">
    <img class="w-full h-auto" src="images/icons/icon-arrow-down.svg" alt="Down Arrow" />
  </div>
  <h3 class="font-medium text-3xl md:text-4xl mx-4 text-center leading-tight tracking-tight">
    Example Result
  </h3>
  <div class="flex flex-wrap justify-center items-center mt-5">
    <!-- Monkey Svg Text -->
    <pre
      class="text-sm bg-black text-orangeLight rounded-lg px-3 py-2 md:px-4 md:py-3 m-2 max-w-lg
      h-64 break-all overflow-y-scroll"
    >
      {#if callSvgText}{callSvgText}{/if}
    </pre>
    <!-- Heisenberg Svg Image -->
    <LazyImage classes="h-64 w-auto m-2">
      <img
        slot="content"
        class="h-full w-auto"
        src="images/illustrations-foreground/call-heisenberg-svg.svg"
        alt="SVG Call Result Heisenberg Illustration"
      />
      <img
        slot="placeholder"
        class="h-full w-auto"
        src="images/illustrations-foreground/call-heisenberg-svg-placeholder.svg"
        alt="SVG Call Result Heisenberg Illustration Placeholder"
      />
    </LazyImage>
  </div>
  <div class="divider mt-20 mb-16" />
  <!-- Optional Parameter -->
  <h2
    class="font-bold text-4xl md:text-5xl mx-4 text-center leading-tight tracking-tight mb-5 md:mb-8"
  >
    Optional Paramaters
  </h2>
  <!-- Format -->
  <div class="w-full flex flex-row flex-wrap justify-center items-center my-6 md:my-5">
    <div class="w-full md:w-1/3 flex md:justify-end text-2xl font-bold my-3">
      <code class="bg-cyanLight px-3 py-1 rounded-lg">format</code>
      <span class="mx-2 md:mx-3">:</span>
    </div>
    <div class="w-full md:w-2/3 md:max-w-lg text-lg leading-loose">
      <code class="font-bold bg-black text-cyanLight rounded-md px-2 mr-1">svg</code>
      (default),
      <code class="font-bold bg-black text-cyanLight rounded-md px-2 mx-1">png</code>
      ,
      <code class="font-bold bg-black text-cyanLight rounded-md px-2 mx-1">webp</code>
      .
    </div>
  </div>
  <!-- Size -->
  <div class="w-full flex flex-row flex-wrap justify-center items-center my-6 md:my-5">
    <div class="w-full md:w-1/3 flex md:justify-end text-2xl font-bold my-3">
      <code class="bg-cyanLight px-3 py-1 rounded-lg">size</code>
      <span class="mx-2 md:mx-3">:</span>
    </div>
    <div class="w-full md:w-2/3 md:max-w-lg text-lg leading-loose text-0">
      <span class="text-lg">In pixels. Ignored when format is</span>
      <code class="text-lg font-bold bg-black text-cyanLight rounded-md px-2 mx-1">svg</code>
      <span class="text-lg">. Default is</span>
      <code class="text-lg font-bold bg-black text-cyanLight rounded-md px-2 mx-1">128</code>
      <span class="text-lg">when format is</span>
      <code class="text-lg font-bold bg-black text-cyanLight rounded-md px-2 mx-1">webp</code>
      <span class="text-lg">or</span>
      <code class="text-lg font-bold bg-black text-cyanLight rounded-md px-2 mx-1">png</code>
      <span class="text-lg">. Minimum is</span>
      <code class="text-lg font-bold bg-black text-cyanLight rounded-md px-2 mx-1">100</code>
      <span class="text-lg">, maximum is</span>
      <code class="text-lg font-bold bg-black text-cyanLight rounded-md px-2 mx-1">1000</code>
      <span class="text-lg">.</span>
    </div>
  </div>
  <!-- Background -->
  <div class="w-full flex flex-row flex-wrap justify-center items-center my-6 md:my-5">
    <div class="w-full md:w-1/3 flex md:justify-end text-2xl font-bold my-3">
      <code class="bg-cyanLight px-3 py-1 rounded-lg">background</code>
      <span class="mx-2 md:mx-3">:</span>
    </div>
    <div class="w-full md:w-2/3 md:max-w-lg text-lg leading-loose">
      <span class="text-lg">Adds a solid color background based on the MonKey.</span>
      <code class="font-bold bg-black text-cyanLight rounded-md px-2 mx-1">false</code>
      <span class="text-lg">(default) or</span>
      <code class="font-bold bg-black text-cyanLight rounded-md px-2 mx-1">true</code>
      <span class="text-lg">.</span>
    </div>
  </div>
  <!-- Arrow Down -->
  <div class="w-12 h-12 my-8">
    <img class="w-full h-auto" src="images/icons/icon-arrow-down.svg" alt="Down Arrow" />
  </div>
  <!-- Example Call -->
  <h3 class="font-medium text-3xl md:text-4xl mx-4 text-center leading-tight tracking-tight">
    Example Call
  </h3>
  <div class="flex flex-wrap justify-center items-center mt-5">
    <code class="text-3xl bg-orangeLight rounded-lg px-3 py-1 font-bold m-2">get</code>
    <div
      class="w-full max-w-2xl flex items-center m-2 text-orangeLight bg-black px-4 py-2 rounded-lg
      mono"
    >
      <span class="text-0 break-all">
        <span class="text-xl">https://monkey.banano.cc/api/v1/</span>
        <span class="text-xl text-purpleLight">
          ban_1ka1ium4pfue3uxtntqsrib8mumxgazsjf58gidh1xeo5te3whsq8z476goo
        </span>
        <span class="text-xl text-white">?</span>
        <span class="text-xl text-cyanLight">format=png</span>
        <span class="text-xl text-white">&</span>
        <span class="text-xl text-cyanLight">size=512</span>
        <span class="text-xl text-white">&</span>
        <span class="text-xl text-cyanLight">background=true</span>
      </span>
    </div>
  </div>
  <!-- Arrow Down -->
  <div class="w-12 h-12 my-8">
    <img class="w-full h-auto" src="images/icons/icon-arrow-down.svg" alt="Down Arrow" />
  </div>
  <h3 class="font-medium text-3xl md:text-4xl mx-4 text-center leading-tight tracking-tight">
    Example Result
  </h3>
  <div class="w-full flex justify-center mt-6">
    <!-- Heisenber PNG Image Desktop -->
    <LazyImage classes="hidden md:block w-full md:w-5/6 lg:w-2/3 h-auto">
      <img
        slot="content"
        class="w-full h-auto"
        src="images/illustrations-foreground/call-heisenberg-png.svg"
        alt="PNG Call Result Heisenberg Illustration"
      />
      <img
        slot="placeholder"
        class="w-full h-auto"
        src="images/illustrations-foreground/call-heisenberg-png-placeholder.svg"
        alt="PNG Call Result Heisenberg Illustration Placeholder"
      />
    </LazyImage>
    <!-- Heisenber PNG Image Mobile -->
    <LazyImage classes="md:hidden w-full h-auto max-w-sm mx-4">
      <img
        slot="content"
        class="w-full h-auto"
        src="images/illustrations-foreground/call-heisenberg-png-mobile.svg"
        alt="PNG Call Result Heisenberg Illustration"
      />
      <img
        slot="placeholder"
        class="w-full h-auto"
        src="images/illustrations-foreground/call-heisenberg-png-mobile-placeholder.svg"
        alt="PNG Call Result Heisenberg Illustration Placeholder"
      />
    </LazyImage>
  </div>
  <div class="divider my-20 md:my-24" />
</div>

<style>
  .text-0 {
    font-size: 0em;
  }
  code,
  pre,
  .mono {
    font-family: "Overpass Mono", monospace;
  }
  pre {
    white-space: pre-wrap !important; /* Since CSS 2.1 */
    white-space: -moz-pre-wrap !important; /* Mozilla, since 1999 */
    white-space: -pre-wrap !important; /* Opera 4-6 */
    white-space: -o-pre-wrap !important; /* Opera 7 */
    word-wrap: break-word !important; /* Internet Explorer 5.5+ */
  }
  .divider {
    background-color: rgba(0, 0, 0, 0.07);
    height: 2px;
    border-radius: 1px;
    width: 95%;
  }
  @font-face {
    font-family: "Overpass Mono";
    src: url("../../fonts/overpass-mono-bold.woff2") format("woff2"),
      url("../../fonts/overpass-mono-bold.woff") format("woff"),
      url("../../fonts/overpass-mono-bold.ttf") format("truetype");
    font-weight: 700;
    font-style: normal;
    font-display: fallback;
  }

  @font-face {
    font-family: "Overpass Mono";
    src: url("../../fonts/overpass-mono-regular.woff2") format("woff2"),
      url("../../fonts/overpass-mono-regular.woff") format("woff"),
      url("../../fonts/overpass-mono-regular.ttf") format("truetype");
    font-weight: 400;
    font-style: normal;
    font-display: fallback;
  }
</style>
