<script lang="ts">
  import demo from '$lib/assets/tf/demo.png';
  import soldier from '$lib/assets/tf/soldier.png';
  import Competition from '$lib/components/display/Competition.svelte';
  import DivisionTag from '$lib/components/display/DivisionTag.svelte';
  import Monthly from '$lib/components/display/Monthly.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import Input from '$lib/components/input/Input.svelte';
  import Label from '$lib/components/input/Label.svelte';
  import Select from '$lib/components/input/select/Select.svelte';
  import SelectButton from '$lib/components/input/select/SelectButton.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import { divisions } from '$lib/src/divisions.js';
  import { SvelteMap } from 'svelte/reactivity';
  import { Temporal } from 'temporal-polyfill';

  import type { MonthlyInputBody } from '$lib/schema.js';

  let { data } = $props();

  // todo: prizes
  let competitionDivisions: SvelteMap<string, string> = new SvelteMap();
  let competitionClass = $state('Soldier');
  let competitionDate = $state('');
  let competitionTime = $state('');

  let monthly: MonthlyInputBody = $derived({
    competition: {
      class: competitionClass,
      starts_at: `${competitionDate} ${competitionTime}`
    },
    divisions: Array.from(competitionDivisions, ([division, map]) => ({ division, map }))
  });
</script>

{#await data.maps then maps}
  {#if maps}
    <div class="flex flex-col gap-1">
      {#each competitionDivisions as [div, _]}
        <div class="flex items-center gap-2">
          <div class="flex w-24 justify-center">
            <DivisionTag {div} />
          </div>
          <Label label="">
            <Select
              type="text"
              options={maps}
              onsubmit={async (value) => {
                competitionDivisions.set(div, value);
                return null;
              }} />
          </Label>
        </div>
      {/each}
    </div>
  {/if}
{/await}

<Monthly {monthly} />

<Section label="create">
  <div class="flex items-center gap-2">
    <SelectButton
      label=""
      options={[
        { src: soldier, value: 'Soldier' },
        { src: demo, value: 'Demo' }
      ]}
      selected={'Soldier'}
      onsubmit={async (value) => {
        competitionClass = value;
        return null;
      }} />
    <div class="flex flex-col">
      <Label label="start date" width={'max-w-60'}>
        <Input
          type="date"
          onsubmit={async (value) => {
            competitionDate = value;
            return null;
          }} />
      </Label>
      <Label label="start time" width={'max-w-60'}>
        <Input
          type="time"
          onsubmit={async (value) => {
            competitionTime = value;
            return null;
          }} />
      </Label>
    </div>
    <div class="ml-auto">
      <span class="text-primary">timezone: </span>
      <span>{Temporal.Now.timeZoneId()}</span>
    </div>
  </div>

  <Label label="add division">
    <Select
      type="text"
      options={divisions}
      onsubmit={async (value) => {
        competitionDivisions.set(value, '');
        return null;
      }} />
  </Label>
</Section>
