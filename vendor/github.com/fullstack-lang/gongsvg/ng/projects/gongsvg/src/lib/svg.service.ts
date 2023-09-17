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

import { SVGDB } from './svg-db';

// insertion point for imports
import { RectDB } from './rect-db'

@Injectable({
  providedIn: 'root'
})
export class SVGService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  SVGServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private svgsUrl: string

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
    this.svgsUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/svgs';
  }

  /** GET svgs from the server */
  getSVGs(GONG__StackPath: string): Observable<SVGDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<SVGDB[]>(this.svgsUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched svgs')),
        catchError(this.handleError<SVGDB[]>('getSVGs', []))
      );
  }

  /** GET svg by id. Will 404 if id not found */
  getSVG(id: number, GONG__StackPath: string): Observable<SVGDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.svgsUrl}/${id}`;
    return this.http.get<SVGDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched svg id=${id}`)),
      catchError(this.handleError<SVGDB>(`getSVG id=${id}`))
    );
  }

  /** POST: add a new svg to the server */
  postSVG(svgdb: SVGDB, GONG__StackPath: string): Observable<SVGDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let Layers = svgdb.Layers
    svgdb.Layers = []
    let StartRect = svgdb.StartRect
    svgdb.StartRect = new RectDB
    let EndRect = svgdb.EndRect
    svgdb.EndRect = new RectDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<SVGDB>(this.svgsUrl, svgdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
	      svgdb.Layers = Layers
        // this.log(`posted svgdb id=${svgdb.ID}`)
      }),
      catchError(this.handleError<SVGDB>('postSVG'))
    );
  }

  /** DELETE: delete the svgdb from the server */
  deleteSVG(svgdb: SVGDB | number, GONG__StackPath: string): Observable<SVGDB> {
    const id = typeof svgdb === 'number' ? svgdb : svgdb.ID;
    const url = `${this.svgsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<SVGDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted svgdb id=${id}`)),
      catchError(this.handleError<SVGDB>('deleteSVG'))
    );
  }

  /** PUT: update the svgdb on the server */
  updateSVG(svgdb: SVGDB, GONG__StackPath: string): Observable<SVGDB> {
    const id = typeof svgdb === 'number' ? svgdb : svgdb.ID;
    const url = `${this.svgsUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let Layers = svgdb.Layers
    svgdb.Layers = []
    let StartRect = svgdb.StartRect
    svgdb.StartRect = new RectDB
    let EndRect = svgdb.EndRect
    svgdb.EndRect = new RectDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<SVGDB>(url, svgdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
	      svgdb.Layers = Layers
        // this.log(`updated svgdb id=${svgdb.ID}`)
      }),
      catchError(this.handleError<SVGDB>('updateSVG'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in SVGService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("SVGService" + error); // log to console instead

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