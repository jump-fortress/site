<script lang="ts">
  import demo from '$lib/assets/tf/demo.png';
  import soldier from '$lib/assets/tf/soldier.png';
  import DivisionTag from '$lib/components/display/DivisionTag.svelte';
  import MonthlyHeader from '$lib/components/display/Monthly.svelte';
  import Button from '$lib/components/input/Button.svelte';
  import Input from '$lib/components/input/Input.svelte';
  import Label from '$lib/components/input/Label.svelte';
  import Select from '$lib/components/input/select/Select.svelte';
  import SelectButton from '$lib/components/input/select/SelectButton.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import { createMonthly } from '$lib/src/api.js';
  import { divisions } from '$lib/src/divisions.js';
  import { SvelteMap } from 'svelte/reactivity';
  import { Temporal } from 'temporal-polyfill';

  import type { Monthly } from '$lib/schema.js';

  let { data } = $props();

  const timeZoneId = Temporal.Now.timeZoneId();
  // todo: prizes
  let competitionDivisions: SvelteMap<string, string> = new SvelteMap();
  let competitionClass = $state('Soldier');
  let competitionDate = $state(Temporal.Now.plainDateISO().add({ weeks: 1 }).toString());
  let competitionTime = $state('00:00');

  // with some placeholder values
  let monthly: Monthly = $derived({
    id: 0,
    competition: {
      id: 0,
      class: competitionClass,
      starts_at: `${competitionDate}T${competitionTime}:00Z`,
      ends_at: '',
      created_at: ''
    },
    divisions: Array.from(competitionDivisions, ([division, map]) => ({
      division,
      map,
      id: 0,
      competition_id: 0
    }))
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

<MonthlyHeader {monthly} />

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
      <span>UTC</span>
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

  <Button
    onsubmit={() => {
      const monthlyValue = monthly;
      monthlyValue.competition.ends_at = monthlyValue.competition.starts_at;
      monthlyValue.competition.created_at = monthlyValue.competition.starts_at;
      return createMonthly(monthlyValue);
    }}>submit monthly</Button>
</Section>
