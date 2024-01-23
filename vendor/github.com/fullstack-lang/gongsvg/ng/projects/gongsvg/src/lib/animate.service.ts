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

import { AnimateDB } from './animate-db'
import { Animate, CopyAnimateToAnimateDB } from './animate'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class AnimateService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  AnimateServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private animatesUrl: string

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
    this.animatesUrl = origin + '/api/github.com/fullstack-lang/gongsvg/go/v1/animates';
  }

  /** GET animates from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<AnimateDB[]> {
    return this.getAnimates(GONG__StackPath, frontRepo)
  }
  getAnimates(GONG__StackPath: string, frontRepo: FrontRepo): Observable<AnimateDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<AnimateDB[]>(this.animatesUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<AnimateDB[]>('getAnimates', []))
      );
  }

  /** GET animate by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<AnimateDB> {
    return this.getAnimate(id, GONG__StackPath, frontRepo)
  }
  getAnimate(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<AnimateDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.animatesUrl}/${id}`;
    return this.http.get<AnimateDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched animate id=${id}`)),
      catchError(this.handleError<AnimateDB>(`getAnimate id=${id}`))
    );
  }

  // postFront copy animate to a version with encoded pointers and post to the back
  postFront(animate: Animate, GONG__StackPath: string): Observable<AnimateDB> {
    let animateDB = new AnimateDB
    CopyAnimateToAnimateDB(animate, animateDB)
    const id = typeof animateDB === 'number' ? animateDB : animateDB.ID
    const url = `${this.animatesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<AnimateDB>(url, animateDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<AnimateDB>('postAnimate'))
    );
  }
  
  /** POST: add a new animate to the server */
  post(animatedb: AnimateDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<AnimateDB> {
    return this.postAnimate(animatedb, GONG__StackPath, frontRepo)
  }
  postAnimate(animatedb: AnimateDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<AnimateDB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON)

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<AnimateDB>(this.animatesUrl, animatedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`posted animatedb id=${animatedb.ID}`)
      }),
      catchError(this.handleError<AnimateDB>('postAnimate'))
    );
  }

  /** DELETE: delete the animatedb from the server */
  delete(animatedb: AnimateDB | number, GONG__StackPath: string): Observable<AnimateDB> {
    return this.deleteAnimate(animatedb, GONG__StackPath)
  }
  deleteAnimate(animatedb: AnimateDB | number, GONG__StackPath: string): Observable<AnimateDB> {
    const id = typeof animatedb === 'number' ? animatedb : animatedb.ID;
    const url = `${this.animatesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<AnimateDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted animatedb id=${id}`)),
      catchError(this.handleError<AnimateDB>('deleteAnimate'))
    );
  }

  // updateFront copy animate to a version with encoded pointers and update to the back
  updateFront(animate: Animate, GONG__StackPath: string): Observable<AnimateDB> {
    let animateDB = new AnimateDB
    CopyAnimateToAnimateDB(animate, animateDB)
    const id = typeof animateDB === 'number' ? animateDB : animateDB.ID
    const url = `${this.animatesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<AnimateDB>(url, animateDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<AnimateDB>('updateAnimate'))
    );
  }

  /** PUT: update the animatedb on the server */
  update(animatedb: AnimateDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<AnimateDB> {
    return this.updateAnimate(animatedb, GONG__StackPath, frontRepo)
  }
  updateAnimate(animatedb: AnimateDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<AnimateDB> {
    const id = typeof animatedb === 'number' ? animatedb : animatedb.ID;
    const url = `${this.animatesUrl}/${id}`;

    // insertion point for reset of pointers (to avoid circular JSON)
    // and encoding of pointers

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<AnimateDB>(url, animatedb, httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers
        // this.log(`updated animatedb id=${animatedb.ID}`)
      }),
      catchError(this.handleError<AnimateDB>('updateAnimate'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in AnimateService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("AnimateService" + error); // log to console instead

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
