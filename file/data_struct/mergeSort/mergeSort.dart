mergeSort(List<int> l, int lo, int hi) {
  if (hi - lo < 2) return;
  int mi = (hi + lo) >> 1;
  mergeSort(l, lo, mi);
  mergeSort(l, mi, hi);
  merge(l, lo, mi, hi);
}

merge(List<int> l, int lo, int mi, int hi) {
  List<int> b = l.sublist(lo, mi);
  List<int> c = l.sublist(mi, hi);
  int lb = mi - lo;
  int lc = hi - mi;
  for (int i = lo, j = 0, k = 0; (j < lb) || (k < lc);) {
    //B数组没有越界 (C数组越界 || B首元素小于C首元素)
    if ((j < lb) && ((lc <= k) || (b[j] <= c[k]))) l[i++] = b[j++];
    //C数组没有越界 && (B数组越界 || C首元素小于C首元素)
    if ((k < lc) && ((lb <= j) || (c[k] <= b[j]))) l[i++] = c[k++];
  }
}

main() {
  List<int> l = [21312, 423, 1231, 31, 23, 123, 1, 2312312];
  mergeSort(l, 0, l.length);
  print(l);
}
