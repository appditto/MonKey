<script context="module">
  export async function preload() {
    let baseUrl = "https://monkey.banano.cc";
    const [res, resMonthly] = await Promise.all([
      this.fetch(`${baseUrl}/api/v1/stats`),
      this.fetch(`${baseUrl}/api/v1/stats/monthly`),
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
    description:
      "Welcome to the jungle! Here, Banano addresses are represented by MonKeys. Come join us and see your MonKey!",
    image: "https://monkey.banano.cc/previews/preview.png",
    imageAlt: "Address visualisation for Banano Addresses",
    url: "https://monkey.banano.cc",
  };
</script>

<!-- HTML -->
<Meta {metadata} />
<Hero />
<How />
<Rigged />
<Available />
<OpenSource />
<Integrate />
<Stats {stats} {statsMonthly} />
