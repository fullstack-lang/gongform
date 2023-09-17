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

import { RectLinkLinkDB } from './rectlinklink-db';

// insertion point for imports
import { RectDB } from './rect-db'
import { LinkDB } from './link-db'
import { LayerDB } from './layer-db'

@Injectable({
  providedIn: 'root'
})
export class RectLinkLinkService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  RectLinkLinkServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private rectlinklinksUrl: string

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
    this.rectlinklinksUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/rectlinklinks';
  }

  /** GET rectlinklinks from the server */
  getRectLinkLinks(GONG__StackPath: string): Observable<RectLinkLinkDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<RectLinkLinkDB[]>(this.rectlinklinksUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched rectlinklinks')),
        catchError(this.handleError<RectLinkLinkDB[]>('getRectLinkLinks', []))
      );
  }

  /** GET rectlinklink by id. Will 404 if id not found */
  getRectLinkLink(id: number, GONG__StackPath: string): Observable<RectLinkLinkDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.rectlinklinksUrl}/${id}`;
    return this.http.get<RectLinkLinkDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched rectlinklink id=${id}`)),
      catchError(this.handleError<RectLinkLinkDB>(`getRectLinkLink id=${id}`))
    );
  }

  /** POST: add a new rectlinklink to the server */
  postRectLinkLink(rectlinklinkdb: RectLinkLinkDB, GONG__StackPath: string): Observable<RectLinkLinkDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let Start = rectlinklinkdb.Start
    rectlinklinkdb.Start = new RectDB
    let End = rectlinklinkdb.End
    rectlinklinkdb.End = new LinkDB
    let _Layer_RectLinkLinks_reverse = rectlinklinkdb.Layer_RectLinkLinks_reverse
    rectlinklinkdb.Layer_RectLinkLinks_reverse = new LayerDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<RectLinkLinkDB>(this.rectlinklinksUrl, rectlinklinkdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        rectlinklinkdb.Layer_RectLinkLinks_reverse = _Layer_RectLinkLinks_reverse
        // this.log(`posted rectlinklinkdb id=${rectlinklinkdb.ID}`)
      }),
      catchError(this.handleError<RectLinkLinkDB>('postRectLinkLink'))
    );
  }

  /** DELETE: delete the rectlinklinkdb from the server */
  deleteRectLinkLink(rectlinklinkdb: RectLinkLinkDB | number, GONG__StackPath: string): Observable<RectLinkLinkDB> {
    const id = typeof rectlinklinkdb === 'number' ? rectlinklinkdb : rectlinklinkdb.ID;
    const url = `${this.rectlinklinksUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<RectLinkLinkDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted rectlinklinkdb id=${id}`)),
      catchError(this.handleError<RectLinkLinkDB>('deleteRectLinkLink'))
    );
  }

  /** PUT: update the rectlinklinkdb on the server */
  updateRectLinkLink(rectlinklinkdb: RectLinkLinkDB, GONG__StackPath: string): Observable<RectLinkLinkDB> {
    const id = typeof rectlinklinkdb === 'number' ? rectlinklinkdb : rectlinklinkdb.ID;
    const url = `${this.rectlinklinksUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let Start = rectlinklinkdb.Start
    rectlinklinkdb.Start = new RectDB
    let End = rectlinklinkdb.End
    rectlinklinkdb.End = new LinkDB
    let _Layer_RectLinkLinks_reverse = rectlinklinkdb.Layer_RectLinkLinks_reverse
    rectlinklinkdb.Layer_RectLinkLinks_reverse = new LayerDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<RectLinkLinkDB>(url, rectlinklinkdb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        rectlinklinkdb.Layer_RectLinkLinks_reverse = _Layer_RectLinkLinks_reverse
        // this.log(`updated rectlinklinkdb id=${rectlinklinkdb.ID}`)
      }),
      catchError(this.handleError<RectLinkLinkDB>('updateRectLinkLink'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in RectLinkLinkService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("RectLinkLinkService" + error); // log to console instead

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
