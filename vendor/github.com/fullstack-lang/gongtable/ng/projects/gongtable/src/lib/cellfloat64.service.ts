// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { CellFloat64DB } from './cellfloat64-db'
import { CellFloat64, CopyCellFloat64ToCellFloat64DB } from './cellfloat64'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class CellFloat64Service {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  CellFloat64ServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private cellfloat64sUrl: string

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
    this.cellfloat64sUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/cellfloat64s';
  }

  /** GET cellfloat64s from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellFloat64DB[]> {
    return this.getCellFloat64s(GONG__StackPath, frontRepo)
  }
  getCellFloat64s(GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellFloat64DB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<CellFloat64DB[]>(this.cellfloat64sUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<CellFloat64DB[]>('getCellFloat64s', []))
      );
  }

  /** GET cellfloat64 by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellFloat64DB> {
    return this.getCellFloat64(id, GONG__StackPath, frontRepo)
  }
  getCellFloat64(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellFloat64DB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.cellfloat64sUrl}/${id}`;
    return this.http.get<CellFloat64DB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched cellfloat64 id=${id}`)),
      catchError(this.handleError<CellFloat64DB>(`getCellFloat64 id=${id}`))
    );
  }

  // postFront copy cellfloat64 to a version with encoded pointers and post to the back
  postFront(cellfloat64: CellFloat64, GONG__StackPath: string): Observable<CellFloat64DB> {
    let cellfloat64DB = new CellFloat64DB
    CopyCellFloat64ToCellFloat64DB(cellfloat64, cellfloat64DB)
    const id = typeof cellfloat64DB === 'number' ? cellfloat64DB : cellfloat64DB.ID
    const url = `${this.cellfloat64sUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<CellFloat64DB>(url, cellfloat64DB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<CellFloat64DB>('postCellFloat64'))
    );
  }
  
  /** POST: add a new cellfloat64 to the server */
  post(cellfloat64db: CellFloat64DB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellFloat64DB> {
    return this.postCellFloat64(cellfloat64db, GONG__StackPath, frontRepo)
  }
  postCellFloat64(cellfloat64db: CellFloat64DB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellFloat64DB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<CellFloat64DB>(this.cellfloat64sUrl, cellfloat64db, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted cellfloat64db id=${cellfloat64db.ID}`)
      }),
      catchError(this.handleError<CellFloat64DB>('postCellFloat64'))
    );
  }

  /** DELETE: delete the cellfloat64db from the server */
  delete(cellfloat64db: CellFloat64DB | number, GONG__StackPath: string): Observable<CellFloat64DB> {
    return this.deleteCellFloat64(cellfloat64db, GONG__StackPath)
  }
  deleteCellFloat64(cellfloat64db: CellFloat64DB | number, GONG__StackPath: string): Observable<CellFloat64DB> {
    const id = typeof cellfloat64db === 'number' ? cellfloat64db : cellfloat64db.ID;
    const url = `${this.cellfloat64sUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<CellFloat64DB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted cellfloat64db id=${id}`)),
      catchError(this.handleError<CellFloat64DB>('deleteCellFloat64'))
    );
  }

  // updateFront copy cellfloat64 to a version with encoded pointers and update to the back
  updateFront(cellfloat64: CellFloat64, GONG__StackPath: string): Observable<CellFloat64DB> {
    let cellfloat64DB = new CellFloat64DB
    CopyCellFloat64ToCellFloat64DB(cellfloat64, cellfloat64DB)
    const id = typeof cellfloat64DB === 'number' ? cellfloat64DB : cellfloat64DB.ID
    const url = `${this.cellfloat64sUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<CellFloat64DB>(url, cellfloat64DB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<CellFloat64DB>('updateCellFloat64'))
    );
  }

  /** PUT: update the cellfloat64db on the server */
  update(cellfloat64db: CellFloat64DB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellFloat64DB> {
    return this.updateCellFloat64(cellfloat64db, GONG__StackPath, frontRepo)
  }
  updateCellFloat64(cellfloat64db: CellFloat64DB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<CellFloat64DB> {
    const id = typeof cellfloat64db === 'number' ? cellfloat64db : cellfloat64db.ID;
    const url = `${this.cellfloat64sUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<CellFloat64DB>(url, cellfloat64db, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated cellfloat64db id=${cellfloat64db.ID}`)
      }),
      catchError(this.handleError<CellFloat64DB>('updateCellFloat64'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in CellFloat64Service', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("CellFloat64Service" + error); // log to console instead

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
