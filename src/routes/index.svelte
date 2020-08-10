<script>
  import Hero from "../components/sections/Hero.svelte";
  import How from "../components/sections/How.svelte";
  import Rigged from "../components/sections/Rigged.svelte";
  import Available from "../components/sections/Available.svelte";
  import OpenSource from "../components/sections/OpenSource.svelte";
  import Integrate from "../components/sections/Integrate.svelte";
  import Stats from "../components/sections/Stats.svelte";
</script>

<style global>
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
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
