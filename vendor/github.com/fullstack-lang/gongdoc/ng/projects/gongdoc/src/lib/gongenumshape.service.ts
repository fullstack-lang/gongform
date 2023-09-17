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

import { GongEnumShapeDB } from './gongenumshape-db';

// insertion point for imports
import { PositionDB } from './position-db'
import { ClassdiagramDB } from './classdiagram-db'

@Injectable({
  providedIn: 'root'
})
export class GongEnumShapeService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  GongEnumShapeServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private gongenumshapesUrl: string

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
    this.gongenumshapesUrl = origin + '/api/github.com/fullstack-lang/gongdoc/go/v1/gongenumshapes';
  }

  /** GET gongenumshapes from the server */
  getGongEnumShapes(GONG__StackPath: string): Observable<GongEnumShapeDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<GongEnumShapeDB[]>(this.gongenumshapesUrl, { params: params })
      .pipe(
        tap(),
		// tap(_ => this.log('fetched gongenumshapes')),
        catchError(this.handleError<GongEnumShapeDB[]>('getGongEnumShapes', []))
      );
  }

  /** GET gongenumshape by id. Will 404 if id not found */
  getGongEnumShape(id: number, GONG__StackPath: string): Observable<GongEnumShapeDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.gongenumshapesUrl}/${id}`;
    return this.http.get<GongEnumShapeDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched gongenumshape id=${id}`)),
      catchError(this.handleError<GongEnumShapeDB>(`getGongEnumShape id=${id}`))
    );
  }

  /** POST: add a new gongenumshape to the server */
  postGongEnumShape(gongenumshapedb: GongEnumShapeDB, GONG__StackPath: string): Observable<GongEnumShapeDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let Position = gongenumshapedb.Position
    gongenumshapedb.Position = new PositionDB
    let GongEnumValueEntrys = gongenumshapedb.GongEnumValueEntrys
    gongenumshapedb.GongEnumValueEntrys = []
    let _Classdiagram_GongEnumShapes_reverse = gongenumshapedb.Classdiagram_GongEnumShapes_reverse
    gongenumshapedb.Classdiagram_GongEnumShapes_reverse = new ClassdiagramDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<GongEnumShapeDB>(this.gongenumshapesUrl, gongenumshapedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
	      gongenumshapedb.GongEnumValueEntrys = GongEnumValueEntrys
        gongenumshapedb.Classdiagram_GongEnumShapes_reverse = _Classdiagram_GongEnumShapes_reverse
        // this.log(`posted gongenumshapedb id=${gongenumshapedb.ID}`)
      }),
      catchError(this.handleError<GongEnumShapeDB>('postGongEnumShape'))
    );
  }

  /** DELETE: delete the gongenumshapedb from the server */
  deleteGongEnumShape(gongenumshapedb: GongEnumShapeDB | number, GONG__StackPath: string): Observable<GongEnumShapeDB> {
    const id = typeof gongenumshapedb === 'number' ? gongenumshapedb : gongenumshapedb.ID;
    const url = `${this.gongenumshapesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<GongEnumShapeDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted gongenumshapedb id=${id}`)),
      catchError(this.handleError<GongEnumShapeDB>('deleteGongEnumShape'))
    );
  }

  /** PUT: update the gongenumshapedb on the server */
  updateGongEnumShape(gongenumshapedb: GongEnumShapeDB, GONG__StackPath: string): Observable<GongEnumShapeDB> {
    const id = typeof gongenumshapedb === 'number' ? gongenumshapedb : gongenumshapedb.ID;
    const url = `${this.gongenumshapesUrl}/${id}`;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    let Position = gongenumshapedb.Position
    gongenumshapedb.Position = new PositionDB
    let GongEnumValueEntrys = gongenumshapedb.GongEnumValueEntrys
    gongenumshapedb.GongEnumValueEntrys = []
    let _Classdiagram_GongEnumShapes_reverse = gongenumshapedb.Classdiagram_GongEnumShapes_reverse
    gongenumshapedb.Classdiagram_GongEnumShapes_reverse = new ClassdiagramDB

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<GongEnumShapeDB>(url, gongenumshapedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
	      gongenumshapedb.GongEnumValueEntrys = GongEnumValueEntrys
        gongenumshapedb.Classdiagram_GongEnumShapes_reverse = _Classdiagram_GongEnumShapes_reverse
        // this.log(`updated gongenumshapedb id=${gongenumshapedb.ID}`)
      }),
      catchError(this.handleError<GongEnumShapeDB>('updateGongEnumShape'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in GongEnumShapeService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("GongEnumShapeService" + error); // log to console instead

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
