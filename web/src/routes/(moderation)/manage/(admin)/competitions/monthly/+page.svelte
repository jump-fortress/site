<script lang="ts">
  import DataSection from '$lib/components/DataSection.svelte';
  import InputSelect from '$lib/components/input/InputSelect.svelte';
  import { Temporal } from 'temporal-polyfill';
  import { divisions } from '$lib/divisions';
  import InputAutofill from '$lib/components/input/InputAutofill.svelte';
  import { onMount } from 'svelte';
  import DivisionTag from '$lib/components/DivisionTag.svelte';

  let validTime = $state(true);

  let { data } = $props();

  let mapNames: string[] = $state([]);

  let competitionMaps: string[] = $state([]);

  function addCompetitionMap(map: string) {
    if (!mapNames.includes(map)) {
      return { error: true, message: 'invalid map name' };
    }
    if (competitionMaps.includes(map)) {
      return { error: true, message: 'map already added' };
    }
    competitionMaps.push(map);
    return { error: false, message: '' };
  }

  onMount(async () => {
    const maps = await data.maps;
    if (maps) {
      mapNames = maps.map((m) => m.name);
    }
  });
</script>

<DataSection title={'Create'}>
  <!-- date & time -->
  <div class="flex h-12 items-center gap-2">
    <label
      for="date"
      class="relative mt-2 w-50 border-2 border-jfgray-700 text-nowrap transition-colors focus-within:border-ctp-lavender-50/50 hover:border-ctp-lavender-50/50"
    >
      <span class="absolute -top-1 left-2 bg-jfgray-800 px-1 text-base leading-1">date</span>
      <div class="flex h-10">
        <input
          class="size-full px-1 text-ctp-lavender transition-colors focus:bg-jfgray-900"
          type="date"
          id="date"
          value={Temporal.Now.plainDateISO().toString()}
        />
      </div>
    </label>
    <label
      for="time"
      class="relative mt-2 w-50 border-2 border-jfgray-700 text-nowrap transition-colors focus-within:border-ctp-lavender-50/50 hover:border-ctp-lavender-50/50 has-invalid:border-ctp-red"
    >
      <span class="absolute -top-1 left-2 bg-jfgray-800 px-1 text-base leading-1">time</span>
      <div class="flex h-10">
        <input
          oninput={(e) => {
            const valid = e.currentTarget.checkValidity();
            validTime = valid;
          }}
          class="size-full px-1 text-ctp-lavender transition-colors focus:bg-jfgray-900"
          type="time"
          id="time"
          step="1800"
          value={Temporal.Now.plainTimeISO().round('hour').toString({ smallestUnit: 'minute' })}
        />
      </div>
    </label>
    {#if !validTime}
      <span>please set a time in steps of 30 minutes (:00 or :30)</span>
    {/if}
    <div class="ml-auto">
      <span class="text-ctp-lavender">timezone: </span>
      <span>{Temporal.Now.timeZoneId()}</span>
    </div>
  </div>
  <!-- add divisions & maps -->
  <InputSelect
    label={'class'}
    options={['Soldier', 'Demo']}
    submitOption={async (val: string) => {
      return { error: false, message: '' };
    }}
  />
  {#await data.maps then maps}
    {#if maps}
      <InputAutofill
        label={'add map'}
        options={maps.map((m) => m.name)}
        submitOption={async (val: string) => {
          return addCompetitionMap(val);
        }}
      />
    {/if}
  {/await}

  <!-- select divisions for maps -->
  {#each competitionMaps as cm}
    <div class="flex items-center gap-2">
      <span>{cm}</span>
      {#each divisions.slice(0, -1) as div}
        <DivisionTag {div} />
      {/each}
    </div>
  {/each}
</DataSection>
