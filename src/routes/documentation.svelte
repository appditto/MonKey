<script>
  import LazyImage from "../components/LazyImage.svelte";
  import axios from "axios";
  import { onMount } from "svelte";
  import Meta from "../components/Meta.svelte";
  const metadata = {
    title: "MonKey | Documentation",
    description: "MonKey API Documentation",
    image: "https://monkey.banano.cc/preview.png",
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

<div class="y-container flex flex-col items-center px-4 pt-20">
  <h1
    class="font-bold text-4xl md:text-6xl mx-5 md:mx-12 text-center
    leading-tight tracking-tight mt-5 md:mt-16 lg:mt-20"
  >
    API Documentation
  </h1>
  <p class="text-center mx-5 mt-4 text-2xl">
    We have a single endpoint and just a couple of paramaters.
  </p>
  <div class="divider mt-20 mb-16" />
  <!-- Endpoint -->
  <h2
    class="font-bold text-4xl md:text-5xl mx-4 text-center leading-tight
    tracking-tight"
  >
    Endpoint
  </h2>
  <div class="flex flex-wrap justify-center mt-8">
    <code class="text-3xl bg-orangeLight rounded-lg px-3 py-1 font-bold m-2">
      get
    </code>
    <div
      class="flex items-center m-2 text-orangeLight bg-black px-4 py-2
      rounded-lg mono"
    >
      <span class="text-0 break-all">
        <span class="text-xl">https://monkey.banano.cc/api/v1/</span>
        <span class="text-xl text-purpleLight">{'<address>'}</span>
      </span>
    </div>
  </div>
  <div class="divider mt-20 mb-16" />
  <!-- Example Call -->
  <h3
    class="font-medium text-3xl md:text-4xl mx-4 text-center leading-tight
    tracking-tight"
  >
    Example Call
  </h3>
  <div class="flex flex-wrap justify-center items-center mt-5">
    <code
      class="flex-shrink text-3xl bg-orangeLight rounded-lg px-3 py-1 font-bold
      m-2"
    >
      get
    </code>
    <div
      class="w-full max-w-2xl flex items-center m-2 text-orangeLight bg-black
      px-4 py-2 rounded-lg mono"
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
    <img
      class="w-full h-auto"
      src="images/icons/icon-arrow-down.svg"
      alt="Down Arrow"
    />
  </div>
  <div class="flex flex-wrap justify-center items-center mt-1">
    <!-- Monkey Svg Text -->
    <pre
      class="text-sm bg-black text-orangeLight rounded-lg px-3 py-2 md:px-4
      md:py-3 m-2 max-w-lg h-64 break-all overflow-y-scroll"
    >
      {#if callSvgText}{callSvgText}{/if}
    </pre>
    <!-- Heisenber Svg Image -->
    <LazyImage classes="h-64 w-auto m-2">
      <img
        slot="content"
        class="h-full w-auto"
        src="images/illustrations-foreground/call-heisenberg-svg.svg"
        alt="Call Result Heisenberg Illustration"
      />
      <img
        slot="placeholder"
        class="h-full w-auto"
        src="images/illustrations-foreground/call-heisenberg-svg-placeholder.svg"
        alt="Call Result Heisenberg Illustration Placeholder"
      />
    </LazyImage>
  </div>
  <div class="divider mt-20 mb-16" />
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
    width: 100%;
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
