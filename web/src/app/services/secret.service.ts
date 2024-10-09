import { HttpClient, HttpParams } from '@angular/common/http';
import { computed, Injectable, signal } from '@angular/core';
import { catchError, EMPTY, map, of, startWith, Subject, switchMap } from 'rxjs';
import { Observable } from 'rxjs/internal/Observable';
import { environment } from 'src/environments/environment';
import { takeUntilDestroyed, toObservable } from '@angular/core/rxjs-interop'

@Injectable({
  providedIn: 'root'
})
export class SecretService {
  private asmServiceUrl = environment.asm.service_url;

  private state = signal<State>({
    loaded: false,
    error: null,
    params: {
      page: 0,
      pageSize: 20,
      keyword: ''
    },
    secretGroups: [],
    totalPages: 1
  });

  //selectors
  params$ = computed(() => this.state().params)
  secretGroups$ = computed(() => this.state().secretGroups)
  error$ = computed(() => this.state().error)
  totalPage$ = computed(() => this.state().totalPages)

  //source
  search$ = new Subject<SearchParams>();

  constructor(
    private http: HttpClient,
  ) {
    this.search$.pipe(
      startWith({
        page: 0,
        pageSize: 20,
        keyword: ''
      }),
      switchMap(params => this.getSecrets(params)
        .pipe(catchError((err) => this.handleError(err)))
      ),
      takeUntilDestroyed()
    ).subscribe(searchReponse => {
      this.state.update((state) => ({
        ...state,
        secretGroups: searchReponse?.list
      }))
    })
  }

  public getSecrets(params: SearchParams): Observable<SearchResponse>{
    let httpParams  = new HttpParams()
      .set('keyword', params.keyword)
      .set('page', params.page)
      .set('pageSize', params.pageSize);
    return this.http.get(this.asmServiceUrl+"/secrets", {observe: 'response', params: httpParams}).pipe(
      map(res => res.body as SearchResponse),
    );
  }
  
  private handleError(err: any): Observable<SearchResponse> {
    this.state.update((state) => ({ ...state, error: err }));
    return EMPTY;
  }

}

export interface State {
  params: SearchParams,
  secretGroups: SecretGroup[],
  totalPages: number,
  loaded: boolean,
  error: string | null
}

export interface SearchParams {
  page: number,
  pageSize: number,
  keyword: string
}

export interface SearchResponse {
  list: SecretGroup[]
}

export interface SecretGroup {
  id: number,
  description: string,
  secrets: Secret[]
}

export interface Secret {
  id: number,
  description: string,
  content: string,
  encrypted?: boolean
}