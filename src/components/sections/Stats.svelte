<script>
  import StatsCard from "../StatsCard.svelte";
</script>

<!-- HTML -->
<div class="w-full pt-16 pb-24 md:pb-40 relative z-10">
  <div class="y-container flex flex-col items-center">
    <h2
      class="font-bold text-5xl mx-4 text-center leading-none tracking-tight
      mt-5">
      MonKey Stats
    </h2>
    <div
      class="w-full flex flex-row flex-wrap justify-center mt-8 px-6 md:px-12">
      <StatsCard stat="0.4M" header="We served" footer="MonKeys in total" />
      <StatsCard stat="124K" header="We served" footer="unique MonKeys" />
      <StatsCard stat="43K" header="We served to" footer="unique clients" />
      <StatsCard stat="5" header="Available on" footer="services" />
    </div>
  </div>
</div>
