<script>
  export let header;
  export let stat;
  export let footer;
  export let classes;
</script>

<style>
  .card {
    border-radius: 0.65rem;
    border-width: 2px;
    border-color: #e4e5e9;
    box-shadow: 0rem 0.3rem 0rem 0rem #e4e5e9;
  }
</style>

<div class="w-full md:w-1/2 lg:w-1/4 max-w-sm px-2 py-3">
  <div
    class="bg-white flex flex-col justify-center items-center card pt-6 pb-8 {classes}">
    <p>{header}</p>
    <p class="font-bold text-4xl">{stat}</p>
    <p>{footer}</p>
  </div>
</div>
