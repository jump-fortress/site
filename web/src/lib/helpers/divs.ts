export const divs: string[] = ['Diamond', 'Platinum', 'Gold', 'Silver', 'Bronze', 'Steel', 'Wood'];

// expects divs to have first letter uppercase
export function compareDivisions(a: string, b: string): number {
  let ai = divs.indexOf(a);
  let bi = divs.indexOf(b);
  if (ai === -1) {
    ai = 10;
  }
  if (bi === -1) {
    bi = 10;
  }
  if (ai === bi) return 0;
  return ai > bi ? 1 : -1;
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
