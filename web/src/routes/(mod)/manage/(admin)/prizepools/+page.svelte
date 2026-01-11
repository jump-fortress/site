<script lang="ts">
  import DivisionTag from '$lib/components/display/DivisionTag.svelte';
  import Table from '$lib/components/display/table/Table.svelte';
  import TableCompetition from '$lib/components/display/table/TableCompetition.svelte';
  import TableDate from '$lib/components/display/table/TableDate.svelte';
  import Button from '$lib/components/input/Button.svelte';
  import Input from '$lib/components/input/Input.svelte';
  import Label from '$lib/components/input/Label.svelte';
  import Section from '$lib/components/layout/Section.svelte';
  import { createPrizepool, deletePrizepool, getCompetitionPrizepool } from '$lib/src/api.js';
  import { slide } from 'svelte/transition';
  import { Temporal } from 'temporal-polyfill';

  import type {
    CompetitionPrize,
    components,
    DivisionPrizepoolInputBody,
    Monthly
  } from '$lib/schema';

  let { data } = $props();

  let competitionId = $state(0);
  let competitionDivisions: components['schemas']['CompetitionDivision'][] = $state([]);

  let prizepoolEdit: DivisionPrizepoolInputBody | null = $state(null);
  let edit = $state(false);
  let refresh = $state(false);

  // refresh upon successful prizepool update
  async function handleResponse(pr: Promise<InputResponse>) {
    const r = await pr;
    if (!r?.error) {
      refresh = !refresh;
      edit = false;
    }
  }
</script>

{#if competitionId}
  {#key refresh}
    {#await getCompetitionPrizepool(competitionId)}
      <span>loading prizepool..</span>
    {:then prizepool}
      <Section label={`id ${competitionId} prizepool`}>
        <div transition:slide>
          {#if prizepool}
            <div class="flex flex-col gap-1">
              {#each prizepool as cdpp}
                <div class="flex items-start gap-2">
                  <span class="flex h-10 w-24 items-center justify-center">
                    <DivisionTag div={cdpp.competition_division.division} />
                  </span>
                  <!-- reset division's prizepool if it exists -->
                  {#if cdpp.prizes?.length}
                    <div class="flex flex-col gap-2">
                      <Button
                        onsubmit={() => {
                          const response = deletePrizepool(cdpp.competition_division.id);
                          handleResponse(response);
                          return response;
                        }}>
                        <span class="icon-[mdi--close]"></span>
                        <span>reset prizepool</span>
                      </Button>
                      <div class="w-fit">
                        <Table data={cdpp.prizes}>
                          {#snippet header()}
                            <th class="w-20">placement</th>
                            <th class="w-24">amount</th>
                          {/snippet}
                          {#snippet row(p: CompetitionPrize)}
                            <td>#{p.placement}</td>
                            <td>{p.amount} keys</td>
                          {/snippet}
                        </Table>
                      </div>
                    </div>

                    <!-- otherwise create -->
                  {:else}
                    <Button
                      onsubmit={async () => {
                        prizepoolEdit = {
                          division_id: cdpp.competition_division.id,
                          division_prizepool: {
                            competition_division: cdpp.competition_division,
                            prizes: []
                          }
                        };
                        edit = true;

                        return null;
                      }}>
                      <span class="icon-[mdi--plus]"></span>
                      <span>new prizepool</span>
                    </Button>
                  {/if}
                </div>
              {/each}
            </div>
          {/if}
        </div>
      </Section>
    {/await}
  {/key}
{/if}

{#if edit && prizepoolEdit}
  <div
    class="relative flex w-full max-w-160 flex-col items-center gap-2 rounded-box bg-base-900 p-2">
    <!-- svelte-ignore a11y_consider_explicit_label -->
    <button
      class="absolute top-2 right-2 cursor-pointer"
      onclick={() => {
        edit = false;
      }}>
      <span class="icon-[mdi--close]"></span>
    </button>

    <span class="flex items-center gap-1">
      <DivisionTag div={prizepoolEdit.division_prizepool.competition_division.division} />
    </span>
    <Label label="number of placements">
      <Input
        type="number"
        onsubmit={async (value) => {
          if (prizepoolEdit) {
            prizepoolEdit.division_prizepool.prizes = [];
            for (let i = 1; i <= parseInt(value); i++) {
              prizepoolEdit.division_prizepool.prizes.push({
                amount: 0,
                division_id: prizepoolEdit.division_id,
                id: 0,
                placement: i
              });
            }
          } else {
            return { error: true, message: "couldn't edit" };
          }

          return null;
        }} />
    </Label>

    <div class="grid grid-cols-4 gap-2">
      {#each prizepoolEdit.division_prizepool.prizes as prize}
        <div class="flex items-center gap-2">
          <span class="w-10 text-end">#{prize.placement}</span>
          <Label label="keys">
            <Input
              type="text"
              placeholder={prize.amount.toString()}
              clear={true}
              onsubmit={async (value) => {
                prize.amount = parseInt(value);

                return null;
              }} />
          </Label>
        </div>
      {/each}
    </div>

    {#if (prizepoolEdit.division_prizepool.prizes ?? []).length !== 0}
      <Button
        onsubmit={async () => {
          if (prizepoolEdit) {
            for (const prize of prizepoolEdit.division_prizepool.prizes ?? []) {
              if (prize.amount === 0) {
                return { error: true, message: `#${prize.placement} missing key amount` };
              }
            }
            const response = createPrizepool(prizepoolEdit);
            handleResponse(response);
            return response;
          } else {
            return { error: true, message: "couldn't create" };
          }
        }}>submit</Button>
    {/if}
  </div>
{/if}

<Section label="editable prizepools">
  {#await data.monthlies then monthlies}
    {#if monthlies}
      {@const nowInstant = Temporal.Now.instant()}
      <Table
        data={monthlies.filter(
          ({ competition }) =>
            nowInstant.until(Temporal.Instant.from(competition.starts_at)).total('minutes') > 0
        )}>
        {#snippet header()}
          <th class="w-12">id</th>
          <th></th>
          <th class="w-20">prizepool</th>
          <th class="w-24">visible</th>
          <th class="w-24">start</th>
          <th class="w-24">end</th>
          <th class="w-20">complete</th>
        {/snippet}
        {#snippet row(m: Monthly)}
          <td>{m.competition.id}</td>
          <td
            onclick={() => {
              competitionId = m.competition.id;
              competitionDivisions = m.divisions ?? [];
            }}
            ><TableCompetition competition={m.competition} format="monthly" formatId={m.id} /></td>
          <td>{m.competition.prizepool ?? 0} keys</td>
          <td><TableDate date={m.competition.visible_at} /></td>
          <td><TableDate date={m.competition.starts_at} /></td>
          <td><TableDate date={m.competition.ends_at} /></td>
          <td>{m.competition.complete}</td>
        {/snippet}
      </Table>
    {/if}
  {/await}
</Section>
