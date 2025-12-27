<script lang="ts">
  import rocket from '$lib/assets/static/classes/rocket.png';
  import sticky from '$lib/assets/static/classes/sticky.png';
  import DivisionTag from './DivisionTag.svelte';

  let { small = false }: { small?: boolean } = $props();
</script>

<!--
@component
todo: deprecated

Preview for monthly / motw competitions
- should support 3 or 4 maps
-->
<div
  class="relative flex h-64 w-full justify-between overflow-hidden rounded-2xl text-shadow-sm/50">
  <!-- top left -->

  <li class="absolute z-10 m-2 flex h-12 items-center gap-2">
    <img src={rocket} alt="soldier" class="h-full object-contain" />
    <span class="text-xl">monthly #1</span>
  </li>

  <!-- top right -->
  <li class="absolute right-0 z-10 mr-2 flex h-16 items-center gap-2">
    <div class="flex flex-col items-end gap-1">
      <div class="flex items-center gap-1">
        <span>400 keys</span>
        <span class="icon-[ri--star-line] size-6"></span>
      </div>
      <div class="flex items-center gap-1">
        <span>10/28/25 15:00 UTC</span>
        <span class="icon-[ri--time-line] size-6"></span>
      </div>
    </div>
  </li>

  <!-- left / right maps -->
  <li class="relative h-full basis-2/5 overflow-hidden">
    <img
      src="https://tempusplaza.xyz/map-backgrounds/jump_mirage.jpg"
      alt=""
      draggable="false"
      class="absolute size-full object-cover select-none" />
    <div class="absolute bottom-0 z-10 m-2">
      <div class="flex flex-col gap-1">
        <span class="span-ellipsis">jump_mirage</span>
        <div class="flex gap-2">
          <DivisionTag div="diamond" />
          <DivisionTag div="platinum" />
          <DivisionTag div="gold" />
        </div>
      </div>
    </div>
  </li>
  <li class="relative h-full basis-2/5 overflow-hidden">
    <img
      src="https://tempusplaza.xyz/map-backgrounds/jump_pokus_rc4.jpg"
      alt=""
      draggable="false"
      class="absolute size-full object-cover select-none" />
    <div class="absolute right-0 bottom-0 z-10 m-2">
      <div class="flex flex-col gap-1">
        <span class="span-ellipsis">jump_pokus_rc4</span>
        <div class="flex justify-end gap-2">
          <DivisionTag div="steel" />
          <DivisionTag div="wood" />
        </div>
      </div>
    </div>
  </li>

  <!-- centered maps -->
  <li class="absolute right-0 left-0 flex h-full justify-center">
    <div class="relative w-3/7 overflow-hidden [clip-path:polygon(20%_0,100%_0,80%_100%,0_100%)]">
      <img
        src="https://tempusplaza.xyz/map-backgrounds/jump_jurf2_a1.jpg"
        alt=""
        draggable="false"
        class="absolute size-full object-cover select-none" />
      <div class="absolute right-1/16 bottom-0 z-10 m-2 flex w-full justify-center">
        <div class="flex flex-col items-center gap-1">
          <span class="span-ellipsis">jump_jurf2_a1</span>
          <div class="flex gap-2">
            <DivisionTag div="silver" />
            <DivisionTag div="bronze" />
          </div>
        </div>
      </div>
    </div>
  </li>
</div>
