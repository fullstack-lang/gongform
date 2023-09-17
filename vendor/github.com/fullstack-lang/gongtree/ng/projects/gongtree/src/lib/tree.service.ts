// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { TreeDB } from './tree-db';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class TreeService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  TreeServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private treesUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.treesUrl = origin + '/api/github.com/fullstack-lang/gongtree/go/v1/trees';
  }

  /** GET trees from the server */
  getTrees(GONG__StackPath: string): Observable<TreeDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<TreeDB[]>(this.treesUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched trees')),
        catchError(this.handleError<TreeDB[]>('getTrees', []))
      );
  }

  /** GET tree by id. Will 404 if id not found */
  getTree(id: number, GONG__StackPath: string): Observable<TreeDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.treesUrl}/${id}`;
    return this.http.get<TreeDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched tree id=${id}`)),
      catchError(this.handleError<TreeDB>(`getTree id=${id}`))
    );
  }

  /** POST: add a new tree to the server */
  postTree(treedb: TreeDB, GONG__StackPath: string): Observable<TreeDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let RootNodes = treedb.RootNodes
    treedb.RootNodes = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<TreeDB>(this.treesUrl, treedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
	      treedb.RootNodes = RootNodes
        // this.log(`posted treedb id=${treedb.ID}`)
      }),
      catchError(this.handleError<TreeDB>('postTree'))
    );
  }

  /** DELETE: delete the treedb from the server */
  deleteTree(treedb: TreeDB | number, GONG__StackPath: string): Observable<TreeDB> {
    const id = typeof treedb === 'number' ? treedb : treedb.ID;
    const url = `${this.treesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<TreeDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted treedb id=${id}`)),
      catchError(this.handleError<TreeDB>('deleteTree'))
    );
  }

  /** PUT: update the treedb on the server */
  updateTree(treedb: TreeDB, GONG__StackPath: string): Observable<TreeDB> {
    const id = typeof treedb === 'number' ? treedb : treedb.ID;
    const url = `${this.treesUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let RootNodes = treedb.RootNodes
    treedb.RootNodes = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<TreeDB>(url, treedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
	      treedb.RootNodes = RootNodes
        // this.log(`updated treedb id=${treedb.ID}`)
      }),
      catchError(this.handleError<TreeDB>('updateTree'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in TreeService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("TreeService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
      console.log(message)
  }
}
