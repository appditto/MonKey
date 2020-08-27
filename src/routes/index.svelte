<script context="module">
  export async function preload() {
    let baseUrl = "https://monkey.banano.cc";
    const [res, resMonthly] = await Promise.all([
      this.fetch(`${baseUrl}/api/v1/stats`),
      this.fetch(`${baseUrl}/api/v1/stats/monthly?month=8`),
    ]);
    const [stats, statsMonthly] = await Promise.all([res.json(), resMonthly.json()]);
    return { stats, statsMonthly };
  }
</script>

<script>
  import Hero from "../components/sections/Hero.svelte";
  import How from "../components/sections/How.svelte";
  import Rigged from "../components/sections/Rigged.svelte";
  import Available from "../components/sections/Available.svelte";
  import OpenSource from "../components/sections/OpenSource.svelte";
  import Integrate from "../components/sections/Integrate.svelte";
  import Stats from "../components/sections/Stats.svelte";
  import Meta from "../components/Meta.svelte";
  export let stats;
  export let statsMonthly;
  const metadata = {
    title: "MonKey",
    description: "Address visualisation for Banano.",
    image: "https://monkey.banano.cc/previews/preview.png",
    imageAlt: "Address visualisation for Banano",
    url: "https://monkey.banano.cc",
  };
</script>

<svelte:head />

<!-- HTML -->
<Meta {metadata} />
<Hero />
<How />
<Rigged />
<Available />
<OpenSource />
<Integrate />
<Stats {stats} {statsMonthly} />

<style global>
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
