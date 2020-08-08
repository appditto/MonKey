<script>
  import Hero from "../components/Hero.svelte";
  import How from "../components/How.svelte";
  import Rigged from "../components/Rigged.svelte";
  import Available from "../components/Available.svelte";
  import OpenSource from "../components/OpenSource.svelte";
  import Integrate from "../components/Integrate.svelte";
  import Stats from "../components/Stats.svelte";
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
