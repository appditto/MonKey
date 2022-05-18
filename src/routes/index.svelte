<script>
  import Hero from "../components/sections/Hero.svelte";
  import How from "../components/sections/How.svelte";
  import Rigged from "../components/sections/Rigged.svelte";
  import Available from "../components/sections/Available.svelte";
  import OpenSource from "../components/sections/OpenSource.svelte";
  import Integrate from "../components/sections/Integrate.svelte";
  /* import Stats from "../components/sections/Stats.svelte"; */
  import Meta from "../components/Meta.svelte";
  /* import { onMount } from "svelte"; */

  /* let stats;
  let statsMonthly;

  onMount(async () => {
    let baseUrl = "https://monkey.banano.cc";
    const [res, resMonthly] = await Promise.all([
      fetch(`${baseUrl}/api/v1/stats`),
      fetch(`${baseUrl}/api/v1/stats/monthly`),
    ]);
    const [_stats, _statsMonthly] = await Promise.all([res.json(), resMonthly.json()]);
    stats = _stats;
    statsMonthly = _statsMonthly;
  }); */

  const metadata = {
    title: "MonKey",
    description:
      "Welcome to the jungle! Here, Banano addresses are represented by MonKeys. Come join us and see your MonKey!",
    image: "https://monkey.banano.cc/previews/preview.png",
    imageAlt: "Address visualisation for Banano Addresses",
    url: "https://monkey.banano.cc",
  };
</script>

<Meta {metadata} />
<Hero />
<How />
<Rigged />
<Available />
<OpenSource />
<Integrate />
<!-- {#if stats && statsMonthly}
  <Stats {stats} {statsMonthly} />
{/if} -->
