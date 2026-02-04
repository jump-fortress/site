import type { Player } from '$lib/schema';

export const divs: string[] = ['Diamond', 'Platinum', 'Gold', 'Silver', 'Bronze', 'Steel', 'Wood'];

// expects divs to have first letter uppercase
function compareDivisions(a: string, b: string): number {
  let ai = divs.indexOf(a);
  let bi = divs.indexOf(b);
  if (ai === -1) {
    ai = 10;
  }
  if (bi === -1) {
    bi = 10;
  }
  if (ai === bi) {
    return 0;
  }
  return ai > bi ? 1 : -1;
}

function compareAlphabetical(a: string, b: string) {
  return a.toLowerCase().localeCompare(b.toLowerCase());
}

export function comparePlayers(a: Player, b: Player, player_class: string): number {
  const adiv: string = player_class === 'Soldier' ? (a.soldier_div ?? '') : (a.demo_div ?? '');
  const bdiv: string = player_class === 'Soldier' ? (b.soldier_div ?? '') : (b.demo_div ?? '');
  const cd: number = compareDivisions(adiv, bdiv);
  if (cd === 0) {
    const ca: number = compareAlphabetical(a.alias, b.alias);
    return ca > 0 ? 1 : -1;
  }
  return cd;
}

export function compareBothDivisions(as: string, bs: string, cs: string, ds: string): number {
  const a = compareDivisions(as, bs) <= 0 ? as : bs;
  const b = compareDivisions(cs, ds) <= 0 ? cs : ds;
  let result = compareDivisions(a, b);
  // highest divisions were even, try lower divisions
  if (result === 0) {
    const a_lower = a === as ? bs : as;
    const b_lower = b === cs ? ds : cs;
    result = compareDivisions(a_lower, b_lower);
  }
  // lower divisions were even, use soldier
  if (result === 0) {
    return compareDivisions(as, cs);
  }
  return result;
}
