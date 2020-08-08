<script>
  import Hero from "../components/sections/Hero.svelte";
  import How from "../components/sections/How.svelte";
  import Rigged from "../components/sections/Rigged.svelte";
  import Available from "../components/sections/Available.svelte";
  import OpenSource from "../components/sections/OpenSource.svelte";
  import Integrate from "../components/sections/Integrate.svelte";
  import Stats from "../components/sections/stats.svelte";
</script>

<style global>
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
  .y-container {
    max-width: 1280px;
    margin-left: auto;
    margin-right: auto;
  }
</style>

<svelte:head />

<!-- HTML -->
<Hero />
<How />
<Rigged />
<Available />
<OpenSource />
<Integrate />
<Stats />
