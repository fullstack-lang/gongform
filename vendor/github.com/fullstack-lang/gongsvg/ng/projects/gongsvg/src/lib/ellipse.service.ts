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

import { EllipseDB } from './ellipse-db'
import { Ellipse, CopyEllipseToEllipseDB } from './ellipse'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { AnimateDB } from './animate-db'

@Injectable({
  providedIn: 'root'
})
export class EllipseService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  EllipseServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private ellipsesUrl: string

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
    this.ellipsesUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/ellipses';
  }

  /** GET ellipses from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<EllipseDB[]> {
    return this.getEllipses(GONG__StackPath, frontRepo)
  }
  getEllipses(GONG__StackPath: string, frontRepo: FrontRepo): Observable<EllipseDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<EllipseDB[]>(this.ellipsesUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<EllipseDB[]>('getEllipses', []))
      );
  }

  /** GET ellipse by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<EllipseDB> {
    return this.getEllipse(id, GONG__StackPath, frontRepo)
  }
  getEllipse(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<EllipseDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.ellipsesUrl}/${id}`;
    return this.http.get<EllipseDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched ellipse id=${id}`)),
      catchError(this.handleError<EllipseDB>(`getEllipse id=${id}`))
    );
  }

  // postFront copy ellipse to a version with encoded pointers and post to the back
  postFront(ellipse: Ellipse, GONG__StackPath: string): Observable<EllipseDB> {
    let ellipseDB = new EllipseDB
    CopyEllipseToEllipseDB(ellipse, ellipseDB)
    const id = typeof ellipseDB === 'number' ? ellipseDB : ellipseDB.ID
    const url = `${this.ellipsesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<EllipseDB>(url, ellipseDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<EllipseDB>('postEllipse'))
    );
  }
  
  /** POST: add a new ellipse to the server */
  post(ellipsedb: EllipseDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<EllipseDB> {
    return this.postEllipse(ellipsedb, GONG__StackPath, frontRepo)
  }
  postEllipse(ellipsedb: EllipseDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<EllipseDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)
    ellipsedb.EllipsePointersEncoding.Animates = []
    for (let _animate of ellipsedb.Animates) {
      ellipsedb.EllipsePointersEncoding.Animates.push(_animate.ID)
    }
    ellipsedb.Animates = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<EllipseDB>(this.ellipsesUrl, ellipsedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        ellipsedb.Animates = new Array<AnimateDB>()
        for (let _id of ellipsedb.EllipsePointersEncoding.Animates) {
          let _animate = frontRepo.Animates.get(_id)
          if (_animate != undefined) {
            ellipsedb.Animates.push(_animate!)
          }
        }
        // this.log(`posted ellipsedb id=${ellipsedb.ID}`)
      }),
      catchError(this.handleError<EllipseDB>('postEllipse'))
    );
  }

  /** DELETE: delete the ellipsedb from the server */
  delete(ellipsedb: EllipseDB | number, GONG__StackPath: string): Observable<EllipseDB> {
    return this.deleteEllipse(ellipsedb, GONG__StackPath)
  }
  deleteEllipse(ellipsedb: EllipseDB | number, GONG__StackPath: string): Observable<EllipseDB> {
    const id = typeof ellipsedb === 'number' ? ellipsedb : ellipsedb.ID;
    const url = `${this.ellipsesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<EllipseDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted ellipsedb id=${id}`)),
      catchError(this.handleError<EllipseDB>('deleteEllipse'))
    );
  }

  // updateFront copy ellipse to a version with encoded pointers and update to the back
  updateFront(ellipse: Ellipse, GONG__StackPath: string): Observable<EllipseDB> {
    let ellipseDB = new EllipseDB
    CopyEllipseToEllipseDB(ellipse, ellipseDB)
    const id = typeof ellipseDB === 'number' ? ellipseDB : ellipseDB.ID
    const url = `${this.ellipsesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<EllipseDB>(url, ellipseDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<EllipseDB>('updateEllipse'))
    );
  }

  /** PUT: update the ellipsedb on the server */
  update(ellipsedb: EllipseDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<EllipseDB> {
    return this.updateEllipse(ellipsedb, GONG__StackPath, frontRepo)
  }
  updateEllipse(ellipsedb: EllipseDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<EllipseDB> {
    const id = typeof ellipsedb === 'number' ? ellipsedb : ellipsedb.ID;
    const url = `${this.ellipsesUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers
    ellipsedb.EllipsePointersEncoding.Animates = []
    for (let _animate of ellipsedb.Animates) {
      ellipsedb.EllipsePointersEncoding.Animates.push(_animate.ID)
    }
    ellipsedb.Animates = []

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<EllipseDB>(url, ellipsedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        ellipsedb.Animates = new Array<AnimateDB>()
        for (let _id of ellipsedb.EllipsePointersEncoding.Animates) {
          let _animate = frontRepo.Animates.get(_id)
          if (_animate != undefined) {
            ellipsedb.Animates.push(_animate!)
          }
        }
        // this.log(`updated ellipsedb id=${ellipsedb.ID}`)
      }),
      catchError(this.handleError<EllipseDB>('updateEllipse'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in EllipseService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("EllipseService" + error); // log to console instead

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
